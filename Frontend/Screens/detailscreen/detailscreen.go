package detailscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	
)

var instructionsStyle = lipgloss.NewStyle().
Align(lipgloss.Left). 
Bold(true). 
Foreground(lipgloss.Color("#216EFFFF"))

var detailsBoxStyle = lipgloss.NewStyle(). 
Margin(0,0,2,0). 
Border(lipgloss.NormalBorder(),true). 
BorderForeground(lipgloss.AdaptiveColor{Light: "#000000FF",Dark: "#EA97FFFF"}). 
Width(70)



type Model struct{
	MovieSelected *API.SearchByIDResponse
	allscreens *allscreens.Model
	width int
	height int
}

func InitializeScreen(MovieSelected *API.SearchByIDResponse,allscreens *allscreens.Model)Model{
	return Model{
		MovieSelected: MovieSelected,
		allscreens:allscreens,
	}
}
func (m *Model) Init() (tea.Cmd){
	return nil;
}

func (m *Model) Update(msg tea.Msg)(Model,tea.Cmd){
	var cmd tea.Cmd

	switch msg:=msg.(type){
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return *m,nil
	case tea.KeyMsg:
		switch msg.String() {
		case "P","p":
			m.allscreens.CurrentScreen = allscreens.Result
		case "R","r":
			m.allscreens.CurrentScreen = allscreens.Search
		case "ctrl+c":
			return *m,tea.Quit
		}
		
	}
	cmd = tea.EnterAltScreen
	return *m,cmd
}

func (m *Model)View()string{
	var movieTitle string =detailsBoxStyle.Render("Title : " +m.MovieSelected.Title+" \n")
	instructions := lipgloss.JoinVertical(
		lipgloss.Left,
		instructionsStyle.Render("Press P to Go Back"),
		instructionsStyle.Render("Press R to Reset Search\n"),
	)
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			movieTitle,
			instructions,
		),
	)
}