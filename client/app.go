package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type appState int

const (
	stateIdle appState = iota
	stateWaiting
)

type App struct {
	state    appState
	messages []Message
	input    textinput.Model
	viewport viewport.Model
	width    int
	height   int
	err      error
}

func NewApp() App {
	ti := textinput.New()
	ti.Placeholder = "Message..."
	ti.Focus()
	ti.CharLimit = 0

	return App{
		input:    ti,
		viewport: viewport.New(0, 0),
	}
}

func (a App) Init() tea.Cmd {
	return tea.Batch(fetchHistory(), textinput.Blink)
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.viewport.Width = msg.Width
		a.viewport.Height = msg.Height - 3
		a.input.Width = msg.Width - 4
		a.renderViewport()
		return a, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return a, tea.Quit
		case tea.KeyEnter:
			if a.state == stateWaiting {
				return a, nil
			}
			content := a.input.Value()
			if content == "" {
				return a, nil
			}
			a.input.Reset()
			a.state = stateWaiting
			a.messages = append(a.messages, Message{Role: "user", Content: content})
			a.renderViewport()
			return a, sendMessage(content)
		}

	case msgSent:
		a.state = stateIdle
		a.err = nil
		a.messages = append(a.messages, Message(msg))
		a.renderViewport()
		return a, nil

	case msgHistory:
		a.messages = []Message(msg)
		a.renderViewport()
		return a, nil

	case msgErr:
		a.state = stateIdle
		a.err = msg.err
		return a, nil
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	a.input, cmd = a.input.Update(msg)
	cmds = append(cmds, cmd)
	a.viewport, cmd = a.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return a, tea.Batch(cmds...)
}
