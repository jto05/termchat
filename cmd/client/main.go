// cmd/client/main.go -- entry point for the chat client.
// Takes a username and server address as arguments.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"termchat/tui"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: client <username> <host:port>")
		os.Exit(1)
	}
	username := os.Args[1]
	serverAddr := os.Args[2]

	p := tea.NewProgram(tui.NewApp(username, serverAddr), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
