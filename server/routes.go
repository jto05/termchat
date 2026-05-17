package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func RegisterRoutes(mux *http.ServeMux, store *MessageStore) {
	mux.HandleFunc("POST /message", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		store.Add(Message{Role: "user", Content: req.Content})

		// Echo reply — replace this with your LLM call
		reply := Message{
			Role:    "assistant",
			Content: fmt.Sprintf("You said: %s", strings.TrimSpace(req.Content)),
		}
		store.Add(reply)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reply)
	})

	mux.HandleFunc("GET /history", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.All())
	})
}
