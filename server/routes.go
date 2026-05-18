package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade HTTP connections to Websocket
// connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
RegisterRoutes()
Takes HTTP connections, upgrades them to Websockets and then routes them
through the Hub.

Gorilla supports only one concurrent writer, so mutex is required for
parallel writes.
*/
func RegisterRoutes(mux *http.ServeMux, hub *Hub) {
	// at websocket endpoint
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // no header for now
		if err != nil {
			log.Printf("error: %v", err)
		}
		defer conn.Close()

		// register/unregister client
		client := make(chan Message, 256)
		hub.register <- client
		// unregister if connection ends
		defer func() { hub.unregister <- client }()

		// write to client any messages that are in its channel
		go func() {
			for msg := range client {
				err := conn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					return
				}
			}
		}()

		// listen through client connection and broadcast any messages read
		for {
			var msg Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Printf("error: %v", err)
			}
			hub.Broadcast(msg)
		}
	})
}
