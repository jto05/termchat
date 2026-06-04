/*
tui/app.go

Defines the bubbletea App model and implements the Elm architecture --
Init, Update, and View -- for the chat client.
*/

package tui

import (
	"termchat/internal/hub"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

/*
App
The bubbletea model. Holds all client state including the WebSocket
connection, message history, input box, and terminal dimensions.
*/
type App struct {
	username   string
	messages   []hub.Message
	input      textinput.Model
	viewport   viewport.Model
	width      int
	height     int
	err        error
	serverAddr string
	conn       *websocket.Conn
}

/*
NewApp
Initializes a new App with the given username and server address.
*/
func NewApp(username string, serverAddr string) App {
	ti := textinput.New()
	ti.Placeholder = "Message..."

	// go into focus state
	ti.Focus()
	ti.CharLimit = 0

	return App{
		username:   username,
		serverAddr: serverAddr,
		input:      ti,
		viewport:   viewport.New(0, 0),
	}
}

/*
Init
Kicks off the WebSocket connection, fetches message history,
and starts the cursor blink on startup.
*/
func (a App) Init() tea.Cmd {
	return tea.Batch(
		connect(a.serverAddr, a.username),
		fetchHistory(a.serverAddr),
		textinput.Blink,
	)
}

/*
Update
Handles all incoming messages -- WebSocket events, keypresses, window
resizes -- and returns the updated model and any follow-up commands.
*/
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
		return a, listenForMessages(a.conn)

	// handle chat history
	case msgHistory:
		a.messages = []hub.Message(m)
		a.viewport.SetContent(a.buildHistory())
		a.viewport.GotoBottom()

	// handle input textbox rendering
	case tea.WindowSizeMsg:
		a.width = m.Width
		a.height = m.Height
		a.viewport.Width = m.Width
		a.viewport.Height = m.Height - 3
		a.input.Width = m.Width - 4

	// build message history given new msg received
	case msgReceived:
		a.messages = append(a.messages, hub.Message(m))
		a.viewport.SetContent(a.buildHistory())
		a.viewport.GotoBottom()
		return a, listenForMessages(a.conn)

	// handle key inputs
	case tea.KeyMsg:
		switch m.Type {
		case tea.KeyEnter: // send message here
			content := a.input.Value()
			a.input.Reset()
			return a, sendMessage(a.conn, content)

		case tea.KeyCtrlC: //
			return a, tea.Quit
		}
	}

	return a, cmd
}
