package states

import (
	"ddd-financial-accounting/internal"
	"ddd-financial-accounting/internal/domain/entities"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

type CreateOperationState struct {
	app        *internal.Application
	inputs     []textinput.Model
	focusIndex int
	done       bool
	width      int
	height     int
}

func NewCreateOperationState(app *internal.Application, _ string) State {
	inputs := make([]textinput.Model, 5)
	for i := range inputs {
		inputs[i] = textinput.New()
	}

	inputs[0].Placeholder = "Set operation type from suggestions"
	inputs[1].Placeholder = "Set bank account id"
	inputs[2].Placeholder = "Set amount of money ( > 0 )"
	inputs[3].Placeholder = "Description ( not necessary )"
	inputs[4].Placeholder = "Set category id"

	inputs[0].SetSuggestions([]string{string(entities.Income), string(entities.Expense)})
	inputs[0].ShowSuggestions = true

	for i := range inputs {
		inputs[i].Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
		inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	}

	inputs[0].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	inputs[0].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	inputs[0].Focus()

	return &CreateOperationState{
		app:        app,
		focusIndex: 0,
		inputs:     inputs,
	}
}

func (m *CreateOperationState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown, tea.KeyUp:
			s := msg.String()
			if s == "down" {
				m.focusIndex++
			} else {
				m.focusIndex--
			}
			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) - 1
			}

			for i := 0; i < len(m.inputs); i++ {
				if i == m.focusIndex {
					m.inputs[i].Focus()
					m.inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
					m.inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
					m.inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
				}
			}
		case tea.KeyEsc:
			return WithApp(m.app, NewOperationMenuState)()
		case tea.KeyEnter:
			if m.focusIndex == len(m.inputs)-1 {
				m.done = true
				return m
			}
			m.focusIndex++
			if m.focusIndex > len(m.inputs)-1 {
				m.focusIndex = len(m.inputs) - 1
			}
			for i := 0; i < len(m.inputs); i++ {
				if i == m.focusIndex {
					m.inputs[i].Focus()
				} else {
					m.inputs[i].Blur()
				}
			}
		}
	}

	for i := range m.inputs {
		m.inputs[i], _ = m.inputs[i].Update(msg)
	}

	return m
}

func (m *CreateOperationState) View() string {
	if m.done {
		operationType := m.inputs[0].Value()
		bankAccountId := m.inputs[1].Value()

		amount, err := strconv.ParseFloat(m.inputs[2].Value(), 64)
		if err != nil {
			return "Invalid amount.\n\nPress Esc to go back.\n"
		}

		description := m.inputs[3].Value()
		categoryId := m.inputs[4].Value()

		operation, err := m.app.OperationFacade.Create(operationType, bankAccountId, amount, description, categoryId)
		if err != nil {
			return fmt.Sprintf("Something went wrong with creating operation: %s.\n\nPress Esc to go back.\n", err.Error())
		}

		return fmt.Sprintf(
			"Submitted!\n\nOperation id:  %s\n\nPress Esc to go back.\n",
			operation.ID,
		)
	}

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		BorderForeground(lipgloss.Color("63"))

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("69"))

	cardContent := titleStyle.Render("Add Operation") + "\n\n" +
		"Operation type:\n" + m.inputs[0].View() + "\n\n" +
		"Bank account id:\n" + m.inputs[1].View() + "\n\n" +
		"Amount:\n" + m.inputs[2].View() + "\n\n" +
		"Description:\n" + m.inputs[3].View() + "\n\n" +
		"Category id:\n" + m.inputs[4].View() + "\n\n" +
		"(Enter to submit, Esc to cancel)"

	return cardStyle.Render(cardContent)
}

func (m *CreateOperationState) Resize(width, height int) {
	m.width = width
	m.height = height
}

func (m *CreateOperationState) Width() int {
	return m.width
}

func (m *CreateOperationState) Height() int {
	return m.height
}

func (m *CreateOperationState) GetMainView() IWithBounds {
	return m
}
