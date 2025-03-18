package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a single WebSocket connection.
type Client struct {
	Hub      *Hub            // The Hub this client belongs to.
	Conn     *websocket.Conn // The WebSocket connection.
	Send     chan []byte     // Outgoing messages.
	Username string          // ✅ Added Username field
}

// NewClient creates a new WebSocket client.
func NewClient(hub *Hub, conn *websocket.Conn, username string) *Client {
	return &Client{
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
		Username: username, // ✅ Initialize username
	}
}

// ReadMessages listens for incoming messages from the client.
func (c *Client) ReadMessages() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("❌ Error reading message:", err)
			break
		}

		c.Hub.Broadcast <- message
	}
}

// WriteMessages sends messages from the Hub to the client.
func (c *Client) WriteMessages() {
	defer c.Conn.Close()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("❌ Error writing message:", err)
			break
		}
	}
}

