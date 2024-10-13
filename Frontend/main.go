package main

import (
	"Frontend/API"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	SearchResults []API.Movie      // Movies in the Search Result
	cursor        int              // Cursor to select a movie
	selected      map[int]struct{} // Which movie is marked as selected
	needsSearch   bool             // Flag to prompt for a movie search
	textInput     textinput.Model  // Text input for search term
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter movie title..."
	ti.Focus() // Start with text input focused
	ti.CharLimit = 256
	ti.Width = 30

	return model{
		SearchResults: []API.Movie{},
		selected:      make(map[int]struct{}),
		needsSearch:   true,
		textInput:     ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink // Start blinking cursor for text input
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.needsSearch {
				// Fetch movie data when "enter" is pressed while in search mode
				movie := strings.TrimSpace(m.textInput.Value())
				if movie == "" {
					return m, nil // Don't search if input is empty
				}

				response, err := API.Search(movie)
				if err != nil {
					fmt.Println("Error searching for movie title:", err)
					return m, nil
				}

				m.SearchResults = response.Search
				m.needsSearch = false
				m.textInput.Reset() // Clear text input after search
				return m, nil
			}

			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.SearchResults)-1 {
				m.cursor++
			}
		case "R","r":
			if m.needsSearch==false {
				m.needsSearch = true
				m.cursor = 0;
				return m,cmd;
			}
		}
	
	}

	// Update textInput when in search mode
	if m.needsSearch {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var output string
	if m.needsSearch {
		output += "Enter The Name Of the Movie You want to add to the Wish List:\n\n"
		output += m.textInput.View() + "\n" // Render the text input field
	} else {
		output+="\n";
		for i, movie := range m.SearchResults {
			cursor := " " // Default cursor is empty space
			if i == m.cursor {
				cursor = ">>>" // Mark the current cursor position
			}
			output += fmt.Sprintf("%s %s (%s)\n", cursor, movie.Title, movie.Year)
			
		}
		output+= "\nPress R to reset Search\n"
	}
	output+= "\nPress q to Quit the App ... \n"
	
	return output
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error Starting Program:", err)
		os.Exit(1)
	}
}
