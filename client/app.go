package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	username string
	messages []Message
	input    textinput.Model
	viewport viewport.Model
	width    int
	height   int
	err      error
}

func NewApp(username string) App {
	ti := textinput.New()
	ti.Placeholder = "Message..."
	ti.Focus()
	ti.CharLimit = 0

	return App{
		username: username,
		input:    ti,
		viewport: viewport.New(0, 0),
	}
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return a, nil
}
