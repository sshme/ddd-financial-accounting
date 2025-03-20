package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewOperationsMenuState(app *internal.Application, _ string) State {
	var items []list.Item
	operations, err := app.OperationFacade.GetAll()
	if err != nil {
		return nil
	}

	for _, operation := range operations {
		items = append(items, NewItem(operation.Date.Format("Monday, January 02, 2006 03:04:05 PM"), "View operation details", WithAppAndId(app, operation.ID, NewOperationsMenuState)))
	}

	items = append(items, NewItem("Back", "Return to operation menu", WithApp(app, NewOperationMenuState)))

	mainView := views.NewMainView(items, "Operations")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
