package main

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"Frontend/Screens/detailscreen"
	"Frontend/Screens/resultscreen"
	"Frontend/Screens/searchscreen"
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
BorderForeground(lipgloss.AdaptiveColor{Light: "#000000",Dark: "#3CFF00FF"})


var quitStyle = lipgloss.NewStyle(). 
Bold(true). 
Foreground(lipgloss.Color("#FF0000FF")).
Align(lipgloss.Center)




type model struct {
	screens       allscreens.Model  
	searchScreen  searchscreen.Model
	resultScreen  resultscreen.Model
	detailscreen detailscreen.Model
	results       []API.Movie
	selectedMovie API.SearchByIDResponse
	width int
	height int
}

func initialModel() *model {
	screens:=allscreens.InitializeScreen()
	return &model{
		screens:      screens,
		searchScreen: searchscreen.InitializeScreen(&screens),
	}
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		m.searchScreen.Init(),
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch  msg:=msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	switch m.screens.CurrentScreen{

		case allscreens.Search:
			
			newSearchScreen, screenCmd := m.searchScreen.Update(msg)
			m.searchScreen = newSearchScreen
			cmd = screenCmd 

			if len(m.searchScreen.SearchResults) > 0 {
				m.results = m.searchScreen.SearchResults
				m.screens.SetScreen(allscreens.Result) 
				m.ChangeScreen(m,allscreens.Result)
			}

		case allscreens.Result:
			newResultScreen, screenCmd := m.resultScreen.Update(msg)
			m.resultScreen = newResultScreen
			cmd = screenCmd 
			if m.screens.CurrentScreen != allscreens.Result {
				m.ChangeScreen(m,m.screens.CurrentScreen)
			}
		case allscreens.Detail:
			newDetailsScreen,screenCmd := m.detailscreen.Update(msg)
			m.detailscreen = newDetailsScreen
			cmd = screenCmd 

			if m.screens.CurrentScreen != allscreens.Detail{
				m.ChangeScreen(m,m.screens.CurrentScreen)
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
	var screenStr string
	switch m.screens.CurrentScreen {
	case allscreens.Search:
		screenStr = m.searchScreen.View() + "\n" + quitStyle.Render("Press ctrl+c to quit the App.\n\n")
	case allscreens.Result:
		screenStr = m.resultScreen.View() + "\n" + quitStyle.Render("Press ctrl+c to quit the App.\n\n")
	case allscreens.Detail:
		screenStr = m.detailscreen.View() + "\n" + quitStyle.Render("Press ctrl+c to quit the App.\n\n")
	}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Left,
			lipgloss.Top,
			lipgloss.JoinVertical(
				lipgloss.Left,
				output,
				screenStr,
			),
		)
}

func main() {
	appModel := initialModel()
	p := tea.NewProgram(appModel,tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
		os.Exit(1)
	}
}


func (m *model) ChangeScreen(currm *model, screen allscreens.Screen)tea.Cmd {
	switch screen {
	case allscreens.Search:
		currm.searchScreen = searchscreen.InitializeScreen(&currm.screens)
	case allscreens.Result:
		currm.resultScreen = resultscreen.InitializeScreen(currm.results, &currm.screens, &currm.selectedMovie)
	case allscreens.Detail:
		currm.detailscreen = detailscreen.InitializeScreen(&currm.selectedMovie, &currm.screens)
	}
	return func() tea.Msg {
		return tea.WindowSizeMsg{Width: currm.width+1, Height: currm.height+1}
	}
}



