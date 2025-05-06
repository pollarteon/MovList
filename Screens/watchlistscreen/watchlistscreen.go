package watchlistscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var instructionsStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Bold(true).
	Foreground(lipgloss.Color("#216EFFFF"))

var Watchlist []API.Movie
var idMap = make(map[string]struct{}) // imdbID → present
const watchlistFile = "./db/watchlist.json"
const pageSize = 15

type Model struct {
	watchlist      []API.Movie
	allscreens     *allscreens.Model
	selectedMovie  *API.SearchByIDResponse
	cursor         int
	pagefirstIndex int
	width          int
	height         int
}

func InitializeScreen(allscreen *allscreens.Model, selectedMovie *API.SearchByIDResponse) Model {
	_ = LoadWatchlistFromFile()
	return Model{
		watchlist:      Watchlist,
		allscreens:     allscreen,
		selectedMovie:  selectedMovie,
		pagefirstIndex: 0,
		cursor:         0,
		width:          0,
		height:         0,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return *m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return *m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			pageLastIndex := m.pagefirstIndex + pageSize
			if pageLastIndex > len(m.watchlist) {
				pageLastIndex = len(m.watchlist)
			}
			if m.cursor < pageLastIndex-m.pagefirstIndex-1 {
				m.cursor++
			}
		case "left", "h":
			if m.pagefirstIndex-pageSize >= 0 {
				m.pagefirstIndex -= pageSize
				m.cursor = 0
			}
		case "right", "l":
			if m.pagefirstIndex+pageSize < len(m.watchlist) {
				m.pagefirstIndex += pageSize
				m.cursor = 0
			}
		case "enter":
			globalIndex := m.pagefirstIndex + m.cursor
			toggleWatchedStatus(globalIndex)
			m.watchlist = Watchlist
		case "s", "S":
			globalIndex := m.pagefirstIndex + m.cursor
			m.allscreens.SetScreen(allscreens.Detail)
			imdbId := m.watchlist[globalIndex].IMDbID
			response, err := API.SearchByID(imdbId)
			if err != nil {
				fmt.Println(err)
				return *m, nil
			}
			*(m.selectedMovie) = response
			return *m, cmd
		case "d", "D":
			globalIndex := m.pagefirstIndex + m.cursor
			imdbId := m.watchlist[globalIndex].IMDbID
			removeFromWatchlist(imdbId)
			m.watchlist = Watchlist
			if m.cursor > 0 && m.cursor >= len(m.watchlist)-m.pagefirstIndex {
				m.cursor--
			}
		case "p", "P":
			m.allscreens.SetScreen(allscreens.Search)
			return *m, cmd
		case "r", "R":
			m.allscreens.SetScreen(allscreens.Result)
			return *m, cmd
		}
	}
	return *m, cmd
}

func (m *Model) View() string {
	var list string
	var watchedSymbol string
	title := "Watchlist\n\n"

	endIndex := m.pagefirstIndex + pageSize
	if endIndex > len(m.watchlist) {
		endIndex = len(m.watchlist)
	}
	page := m.watchlist[m.pagefirstIndex:endIndex]

	for i, movie := range page {
		style := lipgloss.NewStyle().Bold(i == m.cursor)
		if i == m.cursor {
			style = style.Foreground(lipgloss.AdaptiveColor{Light: "#D400FFFF", Dark: "#EA97FFFF"}).Bold(true)
		} else {
			style = style.Foreground(lipgloss.Color("#FFFFFF"))
		}
		if movie.Watched {
			watchedSymbol = "✔️"
		} else {
			watchedSymbol = "❌"
		}
		list += style.Render(fmt.Sprintf("%d. %s (%s) %s", m.pagefirstIndex+i+1, movie.MovieTitle, movie.Year, watchedSymbol)) + "\n"
	}

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			list,
			instructionsStyle.Render("[↑↓] - Navigate  [←→] - Prev/Next Page"),
			instructionsStyle.Render("[S] - View Movie Details"),
			instructionsStyle.Render("[D] - Delete Movie"),
			instructionsStyle.Render("[P] - Go to Search Screen"),
			instructionsStyle.Render("[R] - Go to Results Screen"),
		),
	)
}

func AddToWatchlist(movie API.Movie) {
	if _, exists := idMap[movie.IMDbID]; !exists {
		Watchlist = append(Watchlist, movie)
		idMap[movie.IMDbID] = struct{}{}
		_ = SaveWatchlistToFile()
	}
}

func removeFromWatchlist(imdbID string) {
	newWatchlist := make([]API.Movie, 0, len(Watchlist))
	for _, movie := range Watchlist {
		if movie.IMDbID != imdbID {
			newWatchlist = append(newWatchlist, movie)
		}
	}
	Watchlist = newWatchlist
	delete(idMap, imdbID)
	_ = SaveWatchlistToFile()
}

func toggleWatchedStatus(i int) {
	if i >= 0 && i < len(Watchlist) {
		Watchlist[i].Watched = !Watchlist[i].Watched
		_ = SaveWatchlistToFile()
	}
}

func SaveWatchlistToFile() error {
	data, err := json.MarshalIndent(Watchlist, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(watchlistFile, data, 0644)
}

func LoadWatchlistFromFile() error {
	path := filepath.Join(".", watchlistFile)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(data, &Watchlist)
	if err != nil {
		return err
	}
	for _, movie := range Watchlist {
		idMap[movie.IMDbID] = struct{}{}
	}
	return nil
}
