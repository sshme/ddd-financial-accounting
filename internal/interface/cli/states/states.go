package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/interface/cli/views"
	tea "github.com/charmbracelet/bubbletea"
)

type DecoratedStateFunc func(app *internal.Application, id string) State
type StateFunc func() State

func WithApp(app *internal.Application, f DecoratedStateFunc) StateFunc {
	return func() State {
		return f(app, "")
	}
}

func WithAppAndId(app *internal.Application, id string, f DecoratedStateFunc) StateFunc {
	return func() State {
		return f(app, id)
	}
}

type IWithBounds interface {
	Resize(width, height int)
	Width() int
	Height() int
}

type State interface {
	Update(msg tea.Msg) State
	View() string
	GetMainView() IWithBounds
}

type Item struct {
	title, desc string
	nextState   StateFunc
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }
func (i Item) GetNextState() State { return i.nextState() }

func NewItem(title, desc string, nextState StateFunc) Item {
	return Item{title: title, desc: desc, nextState: nextState}
}

type Menu struct {
	mainView *views.MainView
	app      *internal.Application
}

func (m Menu) Update(msg tea.Msg) State {
	m.mainView.MainList, _ = m.mainView.MainList.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := views.DocStyle.GetFrameSize()
		m.mainView.MainList.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.mainView.MainList.SelectedItem().(Item)
			if ok {
				return i.GetNextState()
			}
		}
	}

	return m
}

func (m Menu) View() string {
	return views.DocStyle.Render(m.mainView.View())
}

func (m Menu) Resize(width, height int) {
	m.mainView.MainList.SetSize(width, height)
}

func (m Menu) Width() int {
	return m.mainView.MainList.Width()
}

func (m Menu) Height() int {
	return m.mainView.MainList.Height()
}

func (m Menu) GetMainView() IWithBounds {
	return m
}
