/*
tui/ui.go

Handles all rendering for the chat TUI. Functions here are pure
presentation -- they read App state and return styled strings.
*/

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

/*
renderViewport
Updates the viewport content with the latest message history
and scrolls to the bottom.
*/
func (a *App) renderViewport() {
	a.viewport.SetContent(a.buildHistory())
	a.viewport.GotoBottom()
}

/*
buildHistory
Formats all messages in App.messages into a single styled string
for display in the viewport. Each message is wrapped to the viewport width.
*/
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

/*
renderInput
Returns the styled input box as a string.
*/
func (a App) renderInput() string {
	return inputBoxStyle.Width(a.width - 2).Render(a.input.View())
}

/*
View
Composes the full TUI layout -- viewport on top, input box on bottom.
Returns "loading..." until the terminal size is known.
*/
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
