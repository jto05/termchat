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


## Architecture

## Concurrency Model

## Systems Diagram

```
┌─────────────────────────────────────────────────────┐
│                      SERVER                         │
│                                                     │
│   cmd/server/main.go                                │
│   └── creates Hub, registers routes, listens :8080  │
│                                                     │
│   internal/hub/hub.go                               │
│   ├── clients: map[chan Message]bool                 │
│   ├── register chan   ──┐                           │
│   ├── unregister chan ──┤── Run() select loop        │
│   └── broadcast chan  ──┘                           │
│                                                     │
│   internal/hub/routes.go                            │
│   ├── GET /history  -- serves stored messages        │
│   └── /ws           -- upgrades to WebSocket        │
│       ├── reads username from query param            │
│       ├── registers client with Hub                  │
│       ├── pushes history to client on connect        │
│       ├── write goroutine: chan -- WebSocket (out)   │
│       └── read loop: WebSocket -- hub.Broadcast      │
└───────────────┬─────────────────────────────────────┘
                │ WebSocket (ws://host:port/ws?username=)
                │ HTTP      (http://host:port/history)
       ┌────────┴────────┐
       │                 │
┌──────▼──────┐   ┌──────▼──────┐
│  CLIENT A   │   │  CLIENT B   │
│  (tcc)      │   │  (tcc)      │
│             │   │             │
│  tui/api.go │   │  tui/api.go │
│  ├── connect│   │  ├── connect│
│  ├── send   │   │  ├── send   │
│  ├── listen │   │  ├── listen │
│  └── fetch  │   │  └── fetch  │
│             │   │             │
│  tui/app.go │   │  tui/app.go │
│  bubbletea  │   │  bubbletea  │
│  Init/Update│   │  Init/Update│
│             │   │             │
│  tui/ui.go  │   │  tui/ui.go  │
│  viewport   │   │  viewport   │
│  + input    │   │  + input    │
└─────────────┘   └─────────────┘
```

**Message flow:**
1. Client types a message and presses enter
2. `sendMessage` writes JSON `{content}` over WebSocket
3. Server read loop receives it, stamps username, calls `hub.Broadcast`
4. Hub appends to history, logs, and sends to every client's channel
5. Each client's write goroutine reads from its channel and writes to WebSocket
6. `listenForMessages` on each client receives it, fires `msgReceived`
7. `Update` appends to `App.messages` and re-renders the viewport
