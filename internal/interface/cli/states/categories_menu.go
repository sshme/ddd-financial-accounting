package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
)

func NewCategoriesMenuState(app *internal.Application, _ string) State {
	var items []list.Item
	categories, err := app.CategoryFacade.GetAll()
	if err != nil {
		return nil
	}

	for _, category := range categories {
		items = append(items, NewItem(category.Name, fmt.Sprintf("ID: %s, Type: %s", category.ID, category.GroupType), WithAppAndId(app, category.ID, NewCategoriesMenuState)))
	}

	items = append(items, NewItem("Back", "Return to category menu", WithApp(app, NewCategoryMenuState)))

	mainView := views.NewMainView(items, "Categories")

	return Menu{
		mainView: mainView,
		app:      app,
	}
}
