package main

import "sync"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageStore struct {
	mu       sync.RWMutex
	messages []Message
}

func NewMessageStore() *MessageStore {
	return &MessageStore{}
}

func (s *MessageStore) Add(msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, msg)
}

func (s *MessageStore) All() []Message {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Message, len(s.messages))
	copy(out, s.messages)
	return out
}
