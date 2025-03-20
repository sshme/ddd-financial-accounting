package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type CategoryDeleteState struct {
	app     *internal.Application
	idInput textinput.Model
	width   int
	height  int
}

func NewCategoryDeleteState(app *internal.Application, _ string) State {
	ti := textinput.New()
	ti.Placeholder = "Enter category id"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return &CategoryDeleteState{
		idInput: ti,
		app:     app,
	}
}

func (s *CategoryDeleteState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.idInput.Value() != "" {
				err := s.app.CategoryFacade.Delete(s.idInput.Value())
				if err != nil {
					return WithApp(s.app, NewCategoryMenuState)()
				}
				return WithApp(s.app, NewCategoryMenuState)()
			}
		case tea.KeyEsc:
			return WithApp(s.app, NewCategoryMenuState)()
		}
	}

	s.idInput, _ = s.idInput.Update(msg)
	return s
}

func (s *CategoryDeleteState) View() string {
	return views.DocStyle.Render(
		"Delete Category\n\n" +
			"Id: " + s.idInput.View() + "\n\n" +
			"(Enter to delete, Esc to cancel)")
}

func (s *CategoryDeleteState) Resize(width, height int) {
	s.width = width
	s.height = height
}

func (s *CategoryDeleteState) Width() int {
	return s.width
}

func (s *CategoryDeleteState) Height() int {
	return s.height
}

func (s *CategoryDeleteState) GetMainView() IWithBounds {
	return s
}
