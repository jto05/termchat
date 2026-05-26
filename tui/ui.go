package tui

import "github.com/charmbracelet/lipgloss"

var (
	nameStyle     = lipgloss.NewStyle().Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	inputBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("8"))
)

func (a *App) renderViewport() {
}

func (a App) buildHistory() string {
	return ""
}

func (a App) View() string {
	if a.conn == nil {
		return "connecting..."
	}
	return "connected as " + a.username
}
