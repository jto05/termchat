package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"termchat/tui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: client <username>")
		os.Exit(1)
	}
	username := os.Args[1]

	p := tea.NewProgram(tui.NewApp(username), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
