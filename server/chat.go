package main

import "sync"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Hub struct {
	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Run() {
}

func (h *Hub) Broadcast(msg Message) {
}
