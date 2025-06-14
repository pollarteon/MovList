package resultscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"Frontend/Screens/watchlistscreen"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#D400FFFF", Dark: "#EA97FFFF"}).
	Bold(true)

var instructionsStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Bold(true).
	Foreground(lipgloss.Color("#216EFFFF"))

type Model struct {
	results        []API.Movie
	cursor         int
	selected       map[int]struct{}
	allscreens     *allscreens.Model
	selectedMovie  *API.SearchByIDResponse
	width          int
	height         int
}

func InitializeScreen(results []API.Movie, allscreens *allscreens.Model, selectedMovie *API.SearchByIDResponse) Model {
	return Model{
		results:        results,
		cursor:         0,
		selected:       make(map[int]struct{}),
		allscreens:     allscreens,
		selectedMovie:  selectedMovie,
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
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.results)-1 {
				m.cursor++
			}
		case "ctrl+c":
			return *m, tea.Quit
		case "S", "s", "enter":
			if len(m.results) == 0 {
				return *m, nil
			}
			m.allscreens.SetScreen(allscreens.Detail)
			selectedMovie := m.results[m.cursor]
			ImdbId := selectedMovie.IMDbID
			response, err := API.SearchByID(ImdbId)
			if err != nil {
				fmt.Println("Error calling API:", err)
				return *m, nil
			}
			if response.Title == "" {
				fmt.Println("Invalid response received:", response)
				return *m, nil
			}
			m.results = []API.Movie{}
			*(m.selectedMovie) = response
			return *m, cmd
		case "R", "r":
			m.allscreens.SetScreen(allscreens.Search)
			return *m, cmd
		case "W", "w":
			m.allscreens.SetScreen(allscreens.Watchlist)
			return *m, cmd
		case "A", "a":
			if len(m.results) > 0 {
				watchlistscreen.AddToWatchlist(m.results[m.cursor])
			}
			return *m, cmd
		}
	}
	return *m, cmd
}


func (m *Model) View() string {
	var list string
	var title string
	var instructions string
	title = "Search Results Screen\n\n"

	for i, movie := range m.results {
		cursor := ""
		movieTitle := fmt.Sprintf("%s %s (%s)", cursor, movie.MovieTitle, movie.Year)
		if i == m.cursor {
			cursor = ">"
			movieTitle = cursorStyle.Render(cursor + movieTitle)
		}
		list += fmt.Sprintf("[%d] %s\n", i+1, movieTitle)
	}
	instructions = instructionsStyle.Render(fmt.Sprintln(`
	[S] - View Movie Details
	[R] - Reset Search
	[A] - Add to Watchlist
	[W] - Watchlist
	`))
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			list,
			instructions,
		),
	)
}
