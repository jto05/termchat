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
	return tea.Batch(connect(a.username), textinput.Blink)
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	a.input, cmd = a.input.Update(msg)

	switch m := msg.(type) {

	// set message errors
	case msgErr:
		a.err = m.err

	// set App's websocket connection when connected to
	case msgConnected:
		a.conn = (*websocket.Conn)(m)

	// handle input textbox rendering
	case tea.WindowSizeMsg:
		a.width = m.Width
		a.height = m.Height
		a.viewport.Width = m.Width
		a.viewport.Height = m.Height - inputHeight
		a.input.Width = m.Width - 4

	// handle key inputs
	case tea.KeyMsg:
		switch m.Type {
		case tea.KeyEnter: // send message here
			content := a.input.Value()
			a.input.Reset()
			return a, sendMessage(a.conn, content)

		case tea.KeyCtrlC, tea.KeyEsc: //
			return a, tea.Quit
		}
	}

	return a, cmd
}
