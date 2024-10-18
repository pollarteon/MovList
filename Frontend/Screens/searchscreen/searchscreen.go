package searchscreen

import (
	"Frontend/API"
	"Frontend/UI/custominput"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var searchInputStyle = lipgloss.NewStyle().
Border(lipgloss.NormalBorder(),false,false,false,true).
Padding(1)


type Model struct {
	MovieInput    custominput.Model
	SearchResults []API.Movie

}

func InitializeScreen() Model {
	return Model{
		MovieInput:    custominput.New("Enter Movie Title ..."),
		SearchResults: []API.Movie{},

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

			response, err := API.Search(movie)
			if err != nil {
				fmt.Println("Error fetching movie titles:", err)
				return *m, nil
			}

			m.SearchResults = response.Search
			m.MovieInput.Reset()
		case "ctrl+c":
			return *m,tea.Quit
		}
	}
	m.MovieInput, cmd = m.MovieInput.Update(msg)

	return *m, cmd
}

func (m Model) View() string {
	output := fmt.Sprintf(`
 

%s 

`, searchInputStyle.Render(m.MovieInput.View()))
	return output
}
