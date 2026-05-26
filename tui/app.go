package tui

import (
	"termchat/internal/hub"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

type App struct {
	username string
	messages []hub.Message
	input    textinput.Model
	viewport viewport.Model
	width    int
	height   int
	err      error
	conn     *websocket.Conn
}

func NewApp(username string) App {
	ti := textinput.New()
	ti.Placeholder = "Message..."

	// go into focus state
	ti.Focus()
	ti.CharLimit = 0

	return App{
		username: username,
		input:    ti,
		viewport: viewport.New(0, 0),
	}
}

func (a App) Init() tea.Cmd {
	return connect(a.username)
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	// set App's websocket connection when connected to
	case msgConnected:
		a.conn = (*websocket.Conn)(m)
	}
	return a, nil
}
