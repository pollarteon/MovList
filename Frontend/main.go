
package main

import (
	"Frontend/API"
	"Frontend/Screens/resultscreen" 
	"Frontend/Screens/searchscreen"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	searchScreen searchscreen.Model
	resultScreen resultscreen.Model
	needsSearch  bool
	results []API.Movie
}

func initialModel() *model {
	return &model{
		searchScreen: searchscreen.InitializeScreen(),
		needsSearch:  true,
	}
}

func (m *model) Init() tea.Cmd {
	return m.searchScreen.Init() // Initialize custom text input with blinking cursor
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.needsSearch{
		newSearchScreen,screenCmd:=m.searchScreen.Update(msg)
		m.searchScreen = newSearchScreen
		cmd = screenCmd;

		if len(m.searchScreen.SearchResults) > 0 {
			m.results = m.searchScreen.SearchResults
			m.needsSearch = false

			// Initialize result screen with search results
			m.resultScreen = resultscreen.InitializeScreen(m.results, &m.searchScreen.MovieInput, &m.needsSearch)
		}
	}else{
		newResultScreen, screenCmd := m.resultScreen.Update(msg)
		m.resultScreen = newResultScreen
		cmd = screenCmd

		// Reset to search screen if needsSearch is toggled
		if m.needsSearch {
			m.searchScreen = searchscreen.InitializeScreen() // Reset the search screen
			cmd = m.searchScreen.Init() // Re-initialize the search screen
		}
	}
	return m,cmd
}

func (m *model) View() string {

	output:=fmt.Sprintf(` 

	__  __            _     _     _   
 |  \/  | _____   _| |   (_)___| |_ 
 | |\/| |/ _ \ \ / / |   | / __| __|
 | |  | | (_) \ V /| |___| \__ \ |_ 
 |_|  |_|\___/ \_/ |_____|_|___/\__|
                                    
 `)

	
	if m.needsSearch{
		output=output+m.searchScreen.View()
		return output;
	}else{
		output=output+m.resultScreen.View()
		return output;
	}
	
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
