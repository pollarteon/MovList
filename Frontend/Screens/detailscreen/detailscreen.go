package detailscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

//STYLING

var instructionsStyle = lipgloss.NewStyle().
Align(lipgloss.Left). 
Bold(true). 
Foreground(lipgloss.Color("#216EFFFF"))

var fieldStyle = lipgloss.NewStyle(). 
Padding(0,1). 
Foreground(lipgloss.AdaptiveColor{Light: "#000000FF",Dark: "#90ADFFFF"}).
Width(100).
Bold(true).
Underline(true)

var valueStyle = lipgloss.NewStyle(). 
Foreground(lipgloss.AdaptiveColor{Light: "#155800FF",Dark: "#EA97FFFF"}). 
Bold(true). 
Width(100)


var EnumeratorStyle = lipgloss.NewStyle(). 
Foreground(lipgloss.Color("#FF7B00FF")). 
MarginRight(1). 
Bold(true)


//COMPONENT DEFINITION

type Model struct{
	MovieSelected *API.SearchByIDResponse
	allscreens *allscreens.Model
	detailsTree *tree.Tree
	width int
	height int
}

func InitializeScreen(MovieSelected *API.SearchByIDResponse,allscreens *allscreens.Model)Model{
	t := tree.New().Root(".").
	Child(
		fieldStyle.Render("Movie Name"),
		tree.New().Child(valueStyle.Render(MovieSelected.Title)),
		fieldStyle.Render("Genre"),
		tree.New().Child(valueStyle.Render(MovieSelected.Genre)),
		fieldStyle.Render("Plot"),
		tree.New().Child(valueStyle.Render(MovieSelected.Plot)),
		fieldStyle.Render("Actors"),
		tree.New().Child(valueStyle.Render(MovieSelected.Actors)),
		fieldStyle.Render("Language"),
		tree.New().Child(valueStyle.Render(MovieSelected.Language)),
		fieldStyle.Render("Country"),
		tree.New().Child(valueStyle.Render(MovieSelected.Country)),
	). 
	Enumerator(tree.RoundedEnumerator). 
	EnumeratorStyle(EnumeratorStyle)


	return Model{
		MovieSelected: MovieSelected,
		allscreens:allscreens,
		detailsTree: t,
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
			m.detailsTree.String()+"\n",
			instructions,
		),
	)
}