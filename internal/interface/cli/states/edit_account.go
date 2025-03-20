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

type EditAccountState struct {
	app         *internal.Application
	bankAccount *entities.BankAccount
	inputs      []textinput.Model
	focusIndex  int
	done        bool
	width       int
	height      int
}

func NewEditAccountCreateState(app *internal.Application, id string) State {
	bankAccount, err := app.AccountFacade.GetAccount(id)
	if err != nil {
		return nil
	}

	inputs := []textinput.Model{textinput.New(), textinput.New()}

	inputs[0].SetValue(bankAccount.Name)
	inputs[0].Placeholder = "Set new name"
	inputs[0].Focus()

	inputs[1].SetValue(fmt.Sprintf("%f", bankAccount.Balance))
	inputs[1].Placeholder = "Set new balance"

	for i := range inputs {
		inputs[i].Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
		inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		inputs[i].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	}

	inputs[0].PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	inputs[0].TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &EditAccountState{
		app:         app,
		bankAccount: bankAccount,
		focusIndex:  0,
		inputs:      inputs,
	}
}

func (m *EditAccountState) Update(msg tea.Msg) State {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyDown, tea.KeyUp:
			s := msg.String()
			if s == "tab" || s == "down" {
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
			return WithApp(m.app, NewAccountMenuState)()
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

func (m *EditAccountState) View() string {
	if m.done {
		err := m.app.AccountFacade.UpdateAccountName(m.bankAccount.ID, m.inputs[0].Value())
		if err != nil {
			return "Something went wrong with updating Name.\n\nPress Esc to go back.\n"
		}

		parsedBalance, err := strconv.ParseFloat(m.inputs[1].Value(), 64)
		if err != nil {
			return "Invalid balance.\n\nPress Esc to go back.\n"
		}
		err = m.app.AccountFacade.SetAccountBalance(m.bankAccount.ID, parsedBalance)
		if err != nil {
			return "Something went wrong with updating Balance.\n\nPress Esc to go back.\n"
		}

		return fmt.Sprintf(
			"Submitted!\n\nName:  %s\nBalance: %s\n\nPress Esc to go back.\n",
			m.inputs[0].Value(),
			m.inputs[1].Value(),
		)
	}

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		BorderForeground(lipgloss.Color("63"))

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("69"))

	cardContent := titleStyle.Render("Bank Account Info") + "\n\n" +
		"Id:\n" + m.bankAccount.ID + "\n\n" +
		"Name:\n" + m.inputs[0].View() + "\n\n" +
		"Balance:\n" + m.inputs[1].View() + "\n\n" +
		"(Tab/Shift+Tab to switch fields, Enter to submit, Esc to cancel)"

	return cardStyle.Render(cardContent)
}

func (m *EditAccountState) Resize(width, height int) {
	m.width = width
	m.height = height
}

func (m *EditAccountState) Width() int {
	return m.width
}

func (m *EditAccountState) Height() int {
	return m.height
}

func (m *EditAccountState) GetMainView() IWithBounds {
	return m
}
