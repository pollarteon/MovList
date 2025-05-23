package searchscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"Frontend/UI/custominput"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//STYLING
var searchInputStyle = lipgloss.NewStyle().
Border(lipgloss.NormalBorder(),false,false,true,false).
Padding(0)
var instructionsStyle = lipgloss.NewStyle().
	Align(lipgloss.Left).
	Bold(true).
	Foreground(lipgloss.Color("#216EFFFF"))



//COMPONENT RENDERING
type Model struct {
	MovieInput    custominput.Model
	SearchResults []API.Movie
	allscreens *allscreens.Model
}

func InitializeScreen(allscreen *allscreens.Model) Model {
	return Model{
		MovieInput:    custominput.New("Enter Movie Title ..."),
		SearchResults: []API.Movie{},
		allscreens: allscreen,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.MovieInput.Init()
}

func (m *Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			movie := strings.TrimSpace(m.MovieInput.Value())
			if movie == "" {
				return *m, nil
			}

			response,err := API.Search(movie);
			if err != nil {
				fmt.Println("Error fetching movie titles:", err)
				return *m, nil
			}

			m.SearchResults = response.Search
			m.allscreens.SetScreen(allscreens.Result)
			m.MovieInput.Reset()
		case "ctrl+w":
			
			m.allscreens.SetScreen(allscreens.Watchlist)
			m.MovieInput.Reset()
			return *m, cmd
		case "ctrl+c":
			return *m,tea.Quit
		}
	}
	m.MovieInput, cmd = m.MovieInput.Update(msg)

	return *m, cmd
}

func (m *Model) View() string {
	output := fmt.Sprintf(`
%s 
%s 

`, searchInputStyle.Render(m.MovieInput.View()),instructionsStyle.Render("\n\nPress ctrl+w to go to watchlist."))
	return output
}
