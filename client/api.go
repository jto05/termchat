package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

const serverURL = "http://localhost:8080"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type (
	msgSent    Message
	msgHistory []Message
	msgErr     struct{ err error }
)

func sendMessage(content string) tea.Cmd {
	return func() tea.Msg {
		body, _ := json.Marshal(map[string]string{"content": content})
		resp, err := http.Post(serverURL+"/message", "application/json", bytes.NewReader(body))
		if err != nil {
			return msgErr{fmt.Errorf("send failed: %w", err)}
		}
		defer resp.Body.Close()

		var msg Message
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			return msgErr{fmt.Errorf("decode failed: %w", err)}
		}
		return msgSent(msg)
	}
}

func fetchHistory() tea.Cmd {
	return func() tea.Msg {
		resp, err := http.Get(serverURL + "/history")
		if err != nil {
			return msgErr{fmt.Errorf("could not connect to server: %w", err)}
		}
		defer resp.Body.Close()

		var msgs []Message
		if err := json.NewDecoder(resp.Body).Decode(&msgs); err != nil {
			return msgErr{fmt.Errorf("decode failed: %w", err)}
		}
		return msgHistory(msgs)
	}
}
