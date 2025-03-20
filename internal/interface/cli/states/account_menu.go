package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewAccountMenuState(app *internal.Application, _ string) State {
	items := []list.Item{
		NewItem("View Accounts", "List all bank accounts", WithApp(app, NewAccountsMenuState)),
		NewItem("Add Account", "Create a new bank account", WithApp(app, NewAccountCreateState)),
		NewItem("Delete Account", "Delete an existing bank account", WithApp(app, NewAccountDeleteState)),
		NewItem("Back", "Return to main menu", WithApp(app, NewMainMenuState)),
	}

	mainView := views.NewMainView(items, "Bank Accounts Menu")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
