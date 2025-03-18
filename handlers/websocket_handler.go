package handlers

import (
	"fmt"
	"net/http"

	"chat-system/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (adjust for security in production)
	},
}

// WebSocketHandler upgrades HTTP requests to WebSocket and registers the client.
func WebSocketHandler(hub *ws.Hub, c *gin.Context) {
	// Get username from query parameter
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
		return
	}

	// Create new client with username
	client := ws.NewClient(hub, conn, username)

	// Register the client in the hub
	hub.Register <- client

	fmt.Printf("✅ New WebSocket connection: %s\n", username)

	// Start goroutines for message handling
	go client.ReadMessages()
	go client.WriteMessages()
}

