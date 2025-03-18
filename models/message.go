package models

// Message represents a chat message.
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
