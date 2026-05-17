package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewMessageStore()
	mux := http.NewServeMux()
	RegisterRoutes(mux, store)

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
