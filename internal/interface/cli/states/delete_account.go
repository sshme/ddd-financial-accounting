package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountDeleteState struct {
	app     *internal.Application
	idInput textinput.Model
	width   int
	height  int
}

func NewAccountDeleteState(app *internal.Application, _ string) State {
	ti := textinput.New()
	ti.Placeholder = "Enter account id"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return &AccountDeleteState{
		idInput: ti,
		app:     app,
	}
}

func (s *AccountDeleteState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.idInput.Value() != "" {
				err := s.app.AccountFacade.DeleteAccount(s.idInput.Value())
				if err != nil {
					return WithApp(s.app, NewAccountMenuState)()
				}
				return WithApp(s.app, NewAccountMenuState)()
			}
		case tea.KeyEsc:
			return WithApp(s.app, NewAccountMenuState)()
		}
	}

	s.idInput, _ = s.idInput.Update(msg)
	return s
}

func (s *AccountDeleteState) View() string {
	return views.DocStyle.Render(
		"Delete Account\n\n" +
			"Id: " + s.idInput.View() + "\n\n" +
			"(Enter to delete, Esc to cancel)")
}

func (s *AccountDeleteState) Resize(width, height int) {
	s.width = width
	s.height = height
}

func (s *AccountDeleteState) Width() int {
	return s.width
}

func (s *AccountDeleteState) Height() int {
	return s.height
}

func (s *AccountDeleteState) GetMainView() IWithBounds {
	return s
}
