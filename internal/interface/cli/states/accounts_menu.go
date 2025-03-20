package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewAccountsMenuState(app *internal.Application, _ string) State {
	var items []list.Item
	accounts, err := app.AccountFacade.GetAllAccounts()
	if err != nil {
		return nil
	}

	for _, account := range accounts {
		items = append(items, NewItem(account.Name, "View account details", WithAppAndId(app, account.ID, NewEditAccountCreateState)))
	}

	items = append(items, NewItem("Back", "Return to account menu", WithApp(app, NewAccountMenuState)))

	mainView := views.NewMainView(items, "Select Bank Account")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
