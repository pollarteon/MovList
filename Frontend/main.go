package main

import (
	"Frontend/API"
	"Frontend/Screens/resultscreen" 
	"Frontend/Screens/searchscreen"
	"Frontend/Screens/allscreens"
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//styling

var style = lipgloss.NewStyle().
Bold(true).
Padding(0,0,0,0).
Foreground(lipgloss.AdaptiveColor{Light: "#000000",Dark: "#3CFF00FF"}).
Align(lipgloss.Center).
Border(lipgloss.NormalBorder(),true).
Width(70).
BorderForeground(lipgloss.AdaptiveColor{Light: "#000000",Dark: "#3CFF00FF"}).
Margin(0) 

var quitStyle = lipgloss.NewStyle(). 
Bold(true). 
Foreground(lipgloss.Color("#FF0000FF")).
Align(lipgloss.Center). 
Width(70)



type model struct {
	screens       allscreens.Model  
	searchScreen  searchscreen.Model
	resultScreen  resultscreen.Model
	needsSearch   bool
	results       []API.Movie
}

func initialModel() *model {
	return &model{
		searchScreen: searchscreen.InitializeScreen(),
		needsSearch:  true,
		screens:      allscreens.InitializeScreen(),
	}
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		m.searchScreen.Init(),
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type){
	case tea.WindowSizeMsg:
		return m,cmd;
	}

	switch m.screens.CurrentScreen{

		case allscreens.Search:
			newSearchScreen, screenCmd := m.searchScreen.Update(msg)
			m.searchScreen = newSearchScreen
			cmd = screenCmd

			if len(m.searchScreen.SearchResults) > 0 {
				m.results = m.searchScreen.SearchResults
				m.screens.SetScreen(allscreens.Result) 

				m.resultScreen = resultscreen.InitializeScreen(m.results, &m.searchScreen.MovieInput,&m.screens)
			}

		case allscreens.Result:
			newResultScreen, screenCmd := m.resultScreen.Update(msg)
			m.resultScreen = newResultScreen
			cmd = screenCmd

			
			if m.screens.CurrentScreen == allscreens.Search {
				m.searchScreen = searchscreen.InitializeScreen() 
				cmd = m.searchScreen.Init()
			}

	}
	return m, cmd
}

func (m model) View() string {
	output := fmt.Sprintf(
`
%s  
`,
		style.Render(`
	__  __            _     _     _   
 |  \/  | _____   _| |   (_)___| |_ 
 | |\/| |/ _ \ \ / / |   | / __| __|
 | |  | | (_) \ V /| |___| \__ \ |_ 
 |_|  |_|\___/ \_/ |_____|_|___/\__|
                                    

"Find. Watch. Repeat. Your Ultimate Movie Companion in the Terminal!"

 `),
)

	if m.screens.CurrentScreen == allscreens.Search {
		output += m.searchScreen.View()+"\n\n"+quitStyle.Render("Press ctrl+c to quit the App .\n\n")
		return output
	} else {
		output += m.resultScreen.View()+"\n\n"+quitStyle.Render("Press ctrl+c to quit the App .\n\n")
		return output
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}
