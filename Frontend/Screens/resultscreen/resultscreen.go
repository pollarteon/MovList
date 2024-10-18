package resultscreen

import (
	"Frontend/API"
	"Frontend/Screens/allscreens"
	"Frontend/UI/custominput"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cursorStyle = lipgloss.NewStyle().
Foreground(lipgloss.AdaptiveColor{Light: "#D400FFFF",Dark: "#EA97FFFF"}).
Bold(true)

var instructionsStyle = lipgloss.NewStyle().
Align(lipgloss.Center). 
Bold(true). 
Foreground(lipgloss.Color("#216EFFFF"))





type Model struct{
	results []API.Movie
	cursor int
	selected map[int]struct{}
	movieInput *custominput.Model
	allscreens *allscreens.Model
}

func InitializeScreen(results []API.Movie,movieInput *custominput.Model,allscreens *allscreens.Model) Model{
	return Model{
		results: results,
		cursor: 0,
		selected: make(map[int]struct{}),
		movieInput: movieInput,
		allscreens: allscreens,
	}
}

func(m Model) Init() tea.Cmd{
	return nil;
} 

func (m Model) Update(msg tea.Msg)(Model,tea.Cmd){
	var cmd tea.Cmd

	switch msg:=msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.results)-1 {
				m.cursor++
			}
		case "ctrl+c":
			return m, tea.Quit
		case "R","r":
			m.movieInput.Reset()
			m.allscreens.SetScreen(allscreens.Search)
			return m,cmd
		}
	}
	return m,cmd
}

func (m Model)View()string{
	var output string
	output+="\nSearch Results Screen\n\n"

	for i,movie:=range m.results{
		cursor:=""
		movieTitle:=fmt.Sprintf("%s %s (%s)",cursor,movie.Title,movie.Year)
		if i==m.cursor{
			cursor = ">"
			movieTitle=cursorStyle.Render(cursor+movieTitle)
		}
		output+=fmt.Sprintf("[%d] %s\n",i+1,movieTitle)
	}
	output+=instructionsStyle.Render( "\n Press R to start a new Movie Search\n")

	return output
}