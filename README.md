# termchat

Termchat is a terminal-based application that allows users to register to a
server and communicate to every user registered to that server.

## Installation

Clone the repository and use the Makefile to build from source

```bash
git clone https://github.com/jto05/termchat
make build
```

## Usage

You can host a termchat server on your machine on port 8080 by running the following binary:
```bash
./cmd/tcs
```

To join this server as a user, run the command, replacing user with your username and localhost:8080 with 
the the address of your server in host:port format:
```bash
./cmd/tcc user localhost:8000
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
