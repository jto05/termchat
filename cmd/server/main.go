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
