package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewOperationMenuState(app *internal.Application, _ string) State {
	items := []list.Item{
		NewItem("View Operations", "List all bank accounts", WithApp(app, NewOperationsMenuState)),
		NewItem("Add Operation", "Create a new bank account", WithApp(app, NewCreateOperationState)),
		NewItem("Delete Operation", "Delete an existing bank account", WithApp(app, NewOperationDeleteState)),
		NewItem("Back", "Return to main menu", WithApp(app, NewMainMenuState)),
	}

	mainView := views.NewMainView(items, "Operations Menu")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
