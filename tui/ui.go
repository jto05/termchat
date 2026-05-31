package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	nameStyle     = lipgloss.NewStyle().Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	inputBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("8"))
)

const inputHeight = 3

func (a *App) renderViewport() {
	a.viewport.SetContent(a.buildHistory())
	a.viewport.GotoBottom()
}

func (a App) buildHistory() string {
	var sb strings.Builder
	for _, msg := range a.messages {
		wrapped := lipgloss.NewStyle().
			Width(a.viewport.Width).
			Render(*msg.Content)

		sb.WriteString(
			nameStyle.Render(*msg.Username) + ": " +
				wrapped + "\n",
		)
	}
	return sb.String()
}

func (a App) renderInput() string {
	return inputBoxStyle.Width(a.width - 2).Render(a.input.View())
}

func (a App) View() string {
	if a.width == 0 {
		return "loading..."
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		a.viewport.View(),
		a.renderInput(),
	)
}
