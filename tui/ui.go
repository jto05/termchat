package tui

import "github.com/charmbracelet/lipgloss"

var (
	nameStyle     = lipgloss.NewStyle().Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	inputBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("8"))
)

const inputHeight = 3

func (a *App) renderViewport() {
}

func (a App) buildHistory() string {
	return ""
}

func (a App) renderInput() string {
	return inputBoxStyle.Width(a.width - 2).Render(a.input.View())
}

func (a App) View() string {
	inputLine := inputBoxStyle.
		Width(a.width - 2).
		Render(a.input.View())

	return lipgloss.JoinVertical(
		lipgloss.Left,
		a.viewport.View(),
		inputLine,
	)
}
