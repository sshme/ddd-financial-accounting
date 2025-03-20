package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewMainMenuState(app *internal.Application, _ string) State {
	items := []list.Item{
		NewItem("Accounts", "Manage bank accounts", WithApp(app, NewAccountMenuState)),
		NewItem("Categories", "Manage income and expense categories", WithApp(app, NewCategoryMenuState)),
		NewItem("Operations", "Manage financial operations", WithApp(app, NewOperationMenuState)),
		NewItem("Analytics", "View financial analytics", WithApp(app, NewMainMenuState)),
		NewItem("Import/Export", "Import or export financial data", WithApp(app, NewMainMenuState)),
	}

	mainView := views.NewMainView(items, "Financial Accounting System")

	return Menu{mainView: mainView, app: app}
}
