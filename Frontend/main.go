// File: main.go
package main

import (
	"Frontend/API"
	"Frontend/UI/custominput" // Import the custom input package
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	SearchResults []API.Movie      // Movies in the Search Result
	cursor        int              // Cursor to select a movie
	selected      map[int]struct{} // Which movie is marked as selected
	needsSearch   bool             // Flag to prompt for a movie search
	movieInput     custominput.Model // Custom text input for search term
}

func initialModel() model {
	return model{
		SearchResults: []API.Movie{},
		selected:      make(map[int]struct{}),
		needsSearch:   true,
		movieInput:     custominput.New("Enter movie title..."), // Use custom text input
	}
}

func (m model) Init() tea.Cmd {
	return m.movieInput.Init() // Initialize custom text input with blinking cursor
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.needsSearch {
				// Fetch movie data when "enter" is pressed while in search mode
				movie := strings.TrimSpace(m.movieInput.Value())
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
				m.movieInput.Reset() // Clear text input after search
				return m, nil
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.SearchResults)-1 {
				m.cursor++
			}
		case "R", "r":
			if !m.needsSearch {
				m.needsSearch = true
				m.cursor = 0
				m.SearchResults = nil // Clear search results for reset
				m.movieInput.Reset()
				return m, cmd
			}
		}
	}

	// Update textInput when in search mode
	if m.needsSearch {
		m.movieInput, cmd = m.movieInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var output string
	output += "\n"
	if m.needsSearch {
		output += "Enter the name of the movie you want to add to the Wish List:\n\n"
		output += m.movieInput.View() + "\n" // Render the custom text input field
	} else {
		for i, movie := range m.SearchResults {
			cursor := " " // Default cursor is empty space
			if i == m.cursor {
				cursor = ">>>" // Mark the current cursor position
			}
			output += fmt.Sprintf("%s %s (%s)\n", cursor, movie.Title, movie.Year)
		}
		output += "\nPress R to start a new search.\n"
	}
	output += "\nPress ctrl+c to quit the app.\n"

	return output
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
