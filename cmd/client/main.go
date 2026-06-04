package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"termchat/tui"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: client <username> <serverURL>")
		os.Exit(1)
	}
	username := os.Args[1]
	serverURL := os.Args[2]

	p := tea.NewProgram(tui.NewApp(username, serverURL), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
