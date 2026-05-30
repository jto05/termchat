package tui

import (
	"termchat/internal/hub"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

// TODO: add a config option to replace all dis?
const serverURL = "ws://localhost:8080/ws"

type (
	msgConnected *websocket.Conn
	msgReceived  hub.Message
	msgHistory   []hub.Message
	msgErr       struct{ err error }
)

func connect(username string) tea.Cmd {
	return func() tea.Msg {
		conn, _, err := websocket.DefaultDialer.Dial(
			serverURL+"?username="+username, // add username with initial query
			nil,
		)
		if err != nil {
			return msgErr{err}
		}
		return msgConnected(conn)
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
