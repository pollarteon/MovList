package resultscreen

import (
	"Frontend/API"
	"Frontend/UI/custominput"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{
	results []API.Movie
	cursor int
	selected map[int]struct{}
	movieInput *custominput.Model
	needsSearch *bool
}

func InitializeScreen(results []API.Movie,movieInput *custominput.Model,needsSearch *bool) Model{
	return Model{
		results: results,
		cursor: 0,
		selected: make(map[int]struct{}),
		movieInput: movieInput,
		needsSearch: needsSearch,
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
			*m.needsSearch = true
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
		if i==m.cursor{
			cursor = "> "
		}
		output+=fmt.Sprintf(" %s %s (%s)\n",cursor,movie.Title,movie.Year)
	}
	output+="\n Press R to start a new Movie Search\n"
	output+="\n Press ctrl+c to quit the App\n\n"

	return output
}