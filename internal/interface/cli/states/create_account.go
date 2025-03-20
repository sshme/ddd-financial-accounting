package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AccountCreateState struct {
	app       *internal.Application
	nameInput textinput.Model
	width     int
	height    int
}

func NewAccountCreateState(app *internal.Application, _ string) State {
	ti := textinput.New()
	ti.Placeholder = "Enter account name"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return &AccountCreateState{
		nameInput: ti,
		app:       app,
	}
}

func (s *AccountCreateState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.nameInput.Value() != "" {
				_, err := s.app.AccountFacade.Create(s.nameInput.Value())
				if err != nil {
					return WithApp(s.app, NewAccountMenuState)()
				}
				return WithApp(s.app, NewAccountMenuState)()
			}
		case tea.KeyEsc:
			return WithApp(s.app, NewAccountMenuState)()
		}
	}

	s.nameInput, _ = s.nameInput.Update(msg)
	return s
}

func (s *AccountCreateState) View() string {
	return views.DocStyle.Render(
		"Create New Account\n\n" +
			"Name: " + s.nameInput.View() + "\n\n" +
			"(Enter to save, Esc to cancel)")
}

func (s *AccountCreateState) Resize(width, height int) {
	s.width = width
	s.height = height
}

func (s *AccountCreateState) Width() int {
	return s.width
}

func (s *AccountCreateState) Height() int {
	return s.height
}

func (s *AccountCreateState) GetMainView() IWithBounds {
	return s
}
