package main

import tea "github.com/charmbracelet/bubbletea"

const serverURL = "http://localhost:8080"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type (
	msgReceived Message
	msgHistory  []Message
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
