package tui

import (
	"encoding/json"
	"net/http"

	"termchat/internal/hub"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

type (
	msgConnected *websocket.Conn
	msgReceived  hub.Message
	msgHistory   []hub.Message
	msgErr       struct{ err error }
)

func connect(serverAddr string, username string) tea.Cmd {
	return func() tea.Msg {
		conn, _, err := websocket.DefaultDialer.Dial(
			"ws://"+serverAddr+"/ws?username="+username, // add username with initial query
			nil,
		)
		if err != nil {
			return msgErr{err}
		}
		return msgConnected(conn)
	}
}

func sendMessage(conn *websocket.Conn, content string) tea.Cmd {
	return func() tea.Msg {
		err := conn.WriteJSON(
			map[string]string{"content": content},
		)
		if err != nil {
			return msgErr{err}
		}
		return nil
	}
}

func listenForMessages(conn *websocket.Conn) tea.Cmd {
	return func() tea.Msg {
		var msg hub.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			return msgErr{err}
		}
		return msgReceived(msg)
	}
}

func fetchHistory(serverAddr string) tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get("http://" + serverAddr + "/history")
		if err != nil {
			return msgErr{err}
		}

		var msgs []hub.Message
		err = json.NewDecoder(resp.Body).Decode(&msgs)
		if err != nil {
			return msgErr{err}
		}
		return msgHistory(msgs)
	}
}
