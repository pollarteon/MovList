// File: custominput/textinput.go
package custominput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	TextInput textinput.Model
}

// New creates and initializes a new custom text input model.
func New(placeholder string) Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50
	return Model{
		TextInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles updates for the custom text input.
func (m *Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.TextInput, cmd = m.TextInput.Update(msg)
	return *m, cmd
}

// View renders the custom text input.
func (m Model) View() string {
	return m.TextInput.View()
}

func (m Model) Value() string {
	return m.TextInput.Value()
}

func (m *Model) Reset() {
	m.TextInput.SetValue("")
}
