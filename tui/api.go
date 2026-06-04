/*
tui/api.go

Handles all communication with the server. Contains the WebSocket
connection, message sending, listening, and history fetching.
*/

package tui

import (
	"encoding/json"
	"net/http"

	"termchat/internal/hub"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

/*
bubbletea message types returned by commands.
msgConnected -- WebSocket connection established.
msgReceived  -- a message arrived from the server.
msgHistory   -- initial message history loaded on connect.
msgErr       -- an error occurred in a command.
*/
type (
	msgConnected *websocket.Conn
	msgReceived  hub.Message
	msgHistory   []hub.Message
	msgErr       struct{ err error }
)

/*
connect
Dials the server WebSocket endpoint and returns msgConnected on success.
The username is passed as a query parameter so the server can stamp it
onto all outgoing messages.
*/
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

/*
sendMessage
Writes a message to the server over the WebSocket connection.
Only content is sent -- the server stamps the username.
*/
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

/*
listenForMessages
Reads a single message from the WebSocket and returns it as msgReceived.
Called repeatedly from Update to keep the listen loop running.
*/
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

/*
fetchHistory
Fetches the full message history from the server via HTTP
and returns it as msgHistory.
*/
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
