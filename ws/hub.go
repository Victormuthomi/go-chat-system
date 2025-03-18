package ws

import (
	"fmt"
	"strings" // ✅ Import added
)

// Hub maintains active clients and broadcasts messages.
type Hub struct {
	Clients    map[*Client]bool    // Active clients
	Usernames  map[string]*Client  // ✅ Fixed: Changed from "Username" to "Usernames"
	Broadcast  chan []byte         // Channel for broadcasting messages
	Register   chan *Client        // Channel for new client registrations
	Unregister chan *Client        // Channel for client disconnections
}

// NewHub creates and returns a new Hub.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Usernames:  make(map[string]*Client), // ✅ Fixed: Use "Usernames"
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub loop to manage clients.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.Usernames[client.Username] = client // ✅ Store username
			fmt.Printf("✅ New client connected: %s\n", client.Username)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(h.Usernames, client.Username) // ✅ Remove from username map
				close(client.Send)
				fmt.Printf("❌ Client disconnected: %s\n", client.Username)
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// BroadcastMessage handles public and private messages.
func (h *Hub) BroadcastMessage(sender *Client, message string) {
	if strings.HasPrefix(message, "@") { // Private message
		parts := strings.SplitN(message, " ", 2)
		if len(parts) < 2 {
			sender.Send <- []byte("Invalid private message format. Use @username message")
			return
		}
		targetUsername := parts[0][1:] // Remove '@'
		privateMessage := fmt.Sprintf("[Private] %s: %s", sender.Username, parts[1])

		if targetClient, exists := h.Usernames[targetUsername]; exists {
			targetClient.Send <- []byte(privateMessage)
			sender.Send <- []byte(privateMessage) // Also show sender
		} else {
			sender.Send <- []byte("User not found: " + targetUsername)
		}
	} else { // Public message
		formattedMessage := fmt.Sprintf("%s: %s", sender.Username, message)
		for client := range h.Clients {
			client.Send <- []byte(formattedMessage)
		}
	}
}

