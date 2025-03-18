package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"chat-system/ws"
  "chat-system/handlers"
)

func main() {
	r := gin.Default()
	hub := ws.NewHub()

	// Start WebSocket Hub in a Goroutine
	go hub.Run()

	// WebSocket Endpoint
	r.GET("/ws", func(c *gin.Context) {
		// Get username from query params
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}
    handlers.WebSocketHandler(hub, c)

	
  })

	// Start Server
	port := ":8080"
	fmt.Println("✅ WebSocket Server running on", port)
	if err := r.Run(port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}

