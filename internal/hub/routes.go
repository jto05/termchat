package hub

import (
	"encoding/json"
	"fmt"
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
		// get usenrame from query
		username := r.URL.Query().Get("username")
		if username == "" {
			// throw error if no username provided
			http.Error(w, "username required", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil) // no header for now
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		defer conn.Close()

		// register/unregister client
		client := make(chan Message, 256)
		hub.register <- client
		for _, msg := range hub.Messages() {
			client <- msg
		}
		// unregister if connection ends
		defer func() { hub.unregister <- client }()

		// logging and broadcasting join messages
		joinMsg := fmt.Sprintf("%s joins the chat", username)
		leaveMsg := fmt.Sprintf("%s leaves the chat", username)
		serverName := "Server"
		log.Println(joinMsg)

		hub.Broadcast(
			Message{
				Username: &serverName,
				Content:  &joinMsg,
			},
		)

		defer hub.Broadcast(
			Message{
				Username: &serverName,
				Content:  &leaveMsg,
			},
		)

		defer log.Println(leaveMsg)

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
			_, raw, err := conn.ReadMessage()
			// check error in connection
			if err != nil {
				log.Printf("error: %v", err)
				return
			}

			// check error in json
			err = json.Unmarshal(raw, &msg)
			if err != nil {
				log.Printf("error: %v", err)
				return
			}

			if msg.Content == nil {
				// check error in message format
				log.Printf("invalid format")
				continue
			}
			msg.Username = &username
			hub.Broadcast(msg)
		}
	})

	// at history endpoint call and return hub.Messages()
	mux.HandleFunc("GET /history",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(hub.Messages())
		})
}
