package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"github.com/charmbracelet/bubbles/list"
)

func NewCategoryMenuState(app *internal.Application, _ string) State {
	items := []list.Item{
		NewItem("View Categories", "List all categories", WithApp(app, NewCategoriesMenuState)),
		NewItem("Add Category", "Create a new category", WithApp(app, NewCategoryCreateState)),
		NewItem("Delete Category", "Delete an existing category", WithApp(app, NewCategoryDeleteState)),
		NewItem("Back", "Return to main menu", WithApp(app, NewMainMenuState)),
	}

	mainView := views.NewMainView(items, "Categories Menu")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
