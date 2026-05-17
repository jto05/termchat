package main

import (
	"log"
	"net/http"
)

func main() {
	hub := NewHub()
	go hub.Run()
	mux := http.NewServeMux()
	RegisterRoutes(mux, hub)

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
