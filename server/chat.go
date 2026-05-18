/*
server/chat.go

A pub-sub router that broadcasts messaeges between subscribers
*/

package main

import "log"

/*
Message
A message contains its contents and an associated Username
*/
type Message struct {
	Username *string `json:"username"`
	Content  *string `json:"content"`
}

/*
Hub
The hub that tracks a list of clients, a channel for registering clients,
a channel of for unregistering clients, and a channel for broadcasting;
uses a "chan Message" as a unique identifier for a client. Clients
is a map for easier deletion
*/
type Hub struct {
	clients    map[chan Message]bool
	register   chan chan Message
	unregister chan chan Message
	broadcast  chan Message
}

/*
NewHub()
Allocates space for a new Hub and returns the struct.
*/
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[chan Message]bool),
		register:   make(chan chan Message),
		unregister: make(chan chan Message),
		broadcast:  make(chan Message),
	}
}

/*
Run()
A loop that intercepts messages in each channel to select from
three actions: register, unregister, broadcast.
*/
func (h *Hub) Run() {
	for {
		select {
		// register new user and add them to clients list
		case client := <-h.register:
			h.clients[client] = true

		// unregister new user by removing them from clients list
		case client := <-h.unregister:
			delete(h.clients, client)
			close(client)

		// broadcast messages to all clients
		case msg := <-h.broadcast:
			log.Printf("[%s]: %s", *msg.Username, *msg.Content)
			for client := range h.clients {
				client <- msg
			}
		}
	}
}

/*
Broadcast
Sends message to hub's broadcast channel
*/
func (h *Hub) Broadcast(msg Message) {
	h.broadcast <- msg
}
