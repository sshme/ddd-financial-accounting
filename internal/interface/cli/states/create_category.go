package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type CategoryCreateState struct {
	app       *internal.Application
	nameInput textinput.Model
	width     int
	height    int
}

func NewCategoryCreateState(app *internal.Application, _ string) State {
	ti := textinput.New()
	ti.Placeholder = "Enter category name"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return &CategoryCreateState{
		nameInput: ti,
		app:       app,
	}
}

func (s *CategoryCreateState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.nameInput.Value() != "" {
				_, _, err := s.app.CategoryFacade.Create(s.nameInput.Value())
				if err != nil {
					return WithApp(s.app, NewCategoryMenuState)()
				}
				return WithApp(s.app, NewCategoryMenuState)()
			}
		case tea.KeyEsc:
			return WithApp(s.app, NewCategoryMenuState)()
		}
	}

	s.nameInput, _ = s.nameInput.Update(msg)
	return s
}

func (s *CategoryCreateState) View() string {
	return views.DocStyle.Render(
		"Create New Category\n\n" +
			"Name: " + s.nameInput.View() + "\n\n" +
			"(Enter to save category in both types Income/Expense, Esc to cancel)")
}

func (s *CategoryCreateState) Resize(width, height int) {
	s.width = width
	s.height = height
}

func (s *CategoryCreateState) Width() int {
	return s.width
}

func (s *CategoryCreateState) Height() int {
	return s.height
}

func (s *CategoryCreateState) GetMainView() IWithBounds {
	return s
}
