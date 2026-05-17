# termchat

```
termchat/
├── go.work
├── server/
│   ├── main.go      — starts HTTP server on :8080
│   ├── chat.go      — in-memory MessageStore (thread-safe)
│   └── routes.go    — POST /message, GET /history
└── client/
    ├── main.go      — bubbletea entry point (alt-screen mode)
    ├── app.go       — state machine (Init / Update)
    ├── ui.go        — View() + lipgloss styling
    └── api.go       — HTTP calls + shared Message type
```
