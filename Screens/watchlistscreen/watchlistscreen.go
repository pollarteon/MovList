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

type Model struct {
	watchlist     []API.Movie
	allscreens    *allscreens.Model
	selectedMovie *API.SearchByIDResponse
	cursor        int
	width         int
	height        int
}

func InitializeScreen(allscreen *allscreens.Model,selectedMovie *API.SearchByIDResponse) Model {

	_ = LoadWatchlistFromFile()
	return Model{
		watchlist:     Watchlist,
		allscreens:    allscreen,
		selectedMovie: selectedMovie,
		cursor:        0,
		width:         0,
		height:        0,
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
			if m.cursor < len(m.watchlist)-1 {
				m.cursor++
			}
		case "enter":
			toggleWatchedStatus(m.cursor)
		case "s","S":
			m.allscreens.SetScreen(allscreens.Detail)
			imdbId := m.watchlist[m.cursor].IMDbID
			response,err :=API.SearchByID(imdbId)
			if err != nil {
				fmt.Println(err)
				return *m, nil
			}
			*(m.selectedMovie) =response;
			return *m,cmd;
		case "d", "D":
			imdbId := m.watchlist[m.cursor].IMDbID
			removeFromWatchlist(imdbId)
			if m.cursor >= len(Watchlist) && m.cursor > 0 {
				m.cursor-- 
			}
			m.watchlist = Watchlist
			return *m, cmd
		case "p", "P":
			m.allscreens.SetScreen(allscreens.Search)
			return *m, cmd
		case "r", "R":
			m.allscreens.SetScreen(allscreens.Result)
			return *m,cmd
		}
	}
	return *m, cmd
}

func (m *Model) View() string {
	var list string
	var watchedSymbol string
	title := "Watchlist\n\n"

	for i, movie := range m.watchlist {
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
		list += style.Render(fmt.Sprintf("%s. %s (%s) %s",fmt.Sprintf("%d",i+1), movie.MovieTitle, movie.Year, watchedSymbol)) + "\n"
	}


	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			list,
			instructionsStyle.Render("[P] - Go Back"),
			instructionsStyle.Render("[R] - Search Results "),
			instructionsStyle.Render("[S] - View Movie Details "),
			instructionsStyle.Render("[D] - Delete Movie "),
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


func toggleWatchedStatus(i int ) {
	Watchlist[i].Watched = !Watchlist[i].Watched
	_ = SaveWatchlistToFile()
}


// SaveWatchlistToFile writes the Watchlist to a JSON file
func SaveWatchlistToFile() error {
	data, err := json.MarshalIndent(Watchlist, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(watchlistFile, data, 0644)
}

// LoadWatchlistFromFile reads the watchlist from file and initializes global variables
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