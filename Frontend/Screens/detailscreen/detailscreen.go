package detailscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	
)

var instructionsStyle = lipgloss.NewStyle().
Align(lipgloss.Left). 
Bold(true). 
Foreground(lipgloss.Color("#216EFFFF"))




type Model struct{
	MovieSelected *API.SearchByIDResponse
	allscreens *allscreens.Model
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

	return *m,cmd
}

func (m Model)View()string{
	
	output := fmt.Sprintf(`

Title:"%s"

%s
%s
`,m.MovieSelected.Title,instructionsStyle.Render("Press P to Go Back"),instructionsStyle.Render("Press R to Reset"))
	return output
}