package views

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle        = lipgloss.NewStyle().Padding(1, 1)
	TitleStyle      = lipgloss.NewStyle().Margin(1, 0, 0).Bold(true).Foreground(lipgloss.Color("#0077cc"))
	ItemStyle       = lipgloss.NewStyle().PaddingLeft(4)
	SelectedStyle   = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00cc77"))
	PaginationStyle = lipgloss.NewStyle().PaddingLeft(4)
)

type MainView struct {
	MainList list.Model
}

func (v *MainView) Resize(width, height int) {
	v.MainList.SetSize(width, height)
}

// NewMainView creates a new MainView with the provided items.
func NewMainView(items []list.Item, title string) *MainView {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = SelectedStyle
	delegate.Styles.NormalTitle = ItemStyle

	mainList := list.New(items, delegate, 0, 0)
	mainList.Title = title
	mainList.SetShowStatusBar(false)
	mainList.SetFilteringEnabled(false)
	mainList.Styles.Title = TitleStyle
	mainList.Styles.PaginationStyle = PaginationStyle

	return &MainView{MainList: mainList}
}

func (v *MainView) Update(msg tea.Msg) (list.Model, tea.Cmd) {
	return v.MainList.Update(msg)
}

func (v *MainView) View() string {
	return DocStyle.Render(v.MainList.View())
}

func (v *MainView) GetSelectedItem() list.Item {
	return v.MainList.SelectedItem()
}
