package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type OperationDeleteState struct {
	app     *internal.Application
	idInput textinput.Model
	width   int
	height  int
}

func NewOperationDeleteState(app *internal.Application, _ string) State {
	ti := textinput.New()
	ti.Placeholder = "Enter account id"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 40

	return &OperationDeleteState{
		idInput: ti,
		app:     app,
	}
}

func (o *OperationDeleteState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if o.idInput.Value() != "" {
				err := o.app.OperationFacade.Delete(o.idInput.Value())
				if err != nil {
					return WithApp(o.app, NewOperationMenuState)()
				}
				return WithApp(o.app, NewOperationMenuState)()
			}
		case tea.KeyEsc:
			return WithApp(o.app, NewOperationMenuState)()
		}
	}

	o.idInput, _ = o.idInput.Update(msg)
	return o
}

func (o *OperationDeleteState) View() string {
	return views.DocStyle.Render(
		"Delete Account\n\n" +
			"Id: " + o.idInput.View() + "\n\n" +
			"(Enter to delete, Esc to cancel)")
}

func (o *OperationDeleteState) Resize(width, height int) {
	o.width = width
	o.height = height
}

func (o *OperationDeleteState) Width() int {
	return o.width
}

func (o *OperationDeleteState) Height() int {
	return o.height
}

func (o *OperationDeleteState) GetMainView() IWithBounds {
	return o
}
