package tui

import (
	"termchat/internal/hub"

	tea "github.com/charmbracelet/bubbletea"
)

const serverURL = "ws://localhost:8080/ws"

type (
	msgReceived hub.Message
	msgHistory  []hub.Message
	msgErr      struct{ err error }
)

func connect(username string) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func sendMessage(username, content string) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func fetchHistory() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
