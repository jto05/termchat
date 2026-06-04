// cmd/server/main.go -- entry point for the chat server.
// Starts the hub and listens for WebSocket connections on :8080.
package main

import (
	"log"
	"net/http"

	"termchat/internal/hub"
)

func main() {
	h := hub.NewHub()
	go h.Run()
	mux := http.NewServeMux()
	hub.RegisterRoutes(mux, h)

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
