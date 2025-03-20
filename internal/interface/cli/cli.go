package cli

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/states"
	"ddd-financial-accounting/internal/interface/cli/views"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	state states.State
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := views.DocStyle.GetFrameSize()
		m.state.GetMainView().Resize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	updatedState := m.state.Update(msg)
	if updatedState != m.state {
		updatedState.GetMainView().Resize(
			m.state.GetMainView().Width(),
			m.state.GetMainView().Height())
		m.state = updatedState
	}

	return m, nil
}

func (m Model) View() string {
	return m.state.View()
}

func NewProgram(app *internal.Application) *tea.Program {
	model := Model{
		state: states.WithApp(app, states.NewMainMenuState)(),
	}

	return tea.NewProgram(model, tea.WithAltScreen())
}
