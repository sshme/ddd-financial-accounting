package views

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type FormView struct {
	inputs     []textinput.Model
	formResult string
}

func NewFormView(inputs []textinput.Model, formResult string) *FormView {
	return &FormView{inputs: inputs, formResult: formResult}
}

func (v *FormView) View() string {
	formStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#33aaff")).
		Padding(1, 2).
		Width(60)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#0077cc")).
		MarginBottom(1)

	inputStyle := lipgloss.NewStyle().
		MarginTop(1)

	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#33cc33")).
		MarginTop(1)

	helpStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#999999")).
		MarginTop(1)

	var formContent []string
	formContent = append(formContent, titleStyle.Render("Form"))

	for _, input := range v.inputs {
		formContent = append(formContent, inputStyle.Render(input.View()))
	}

	if v.formResult != "" {
		resultStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cc3333")).
			MarginTop(1)
		formContent = append(formContent, resultStyle.Render(v.formResult))
	}

	formContent = append(formContent,
		buttonStyle.Render("[Enter] Submit"),
		helpStyle.Render("Press ESC to cancel"),
	)

	return formStyle.Render(lipgloss.JoinVertical(lipgloss.Left, formContent...))
}
