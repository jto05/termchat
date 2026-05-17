package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	labelUser = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12")).Render("You")
	labelBot  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Render("Assistant")

	waitingStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	inputBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8"))
)

func (a *App) renderViewport() {
	a.viewport.SetContent(a.buildHistory())
	a.viewport.GotoBottom()
}

func (a App) buildHistory() string {
	var sb strings.Builder
	for _, msg := range a.messages {
		switch msg.Role {
		case "user":
			sb.WriteString(labelUser + "\n")
		default:
			sb.WriteString(labelBot + "\n")
		}
		sb.WriteString(msg.Content + "\n\n")
	}
	if a.state == stateWaiting {
		sb.WriteString(labelBot + "\n")
		sb.WriteString(waitingStyle.Render("...") + "\n")
	}
	return sb.String()
}

func (a App) View() string {
	if a.width == 0 {
		return "Loading..."
	}

	var inputLine string
	if a.err != nil {
		inputLine = errorStyle.Render(fmt.Sprintf("Error: %v", a.err))
	} else {
		inputLine = inputBoxStyle.Width(a.width - 2).Render(a.input.View())
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		a.viewport.View(),
		inputLine,
	)
}
