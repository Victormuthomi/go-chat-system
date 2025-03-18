package routes

import (
	"chat-system/handlers"
	"chat-system/ws"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the HTTP routes.
func RegisterRoutes(router *gin.Engine, hub *ws.Hub) {
	// WebSocket endpoint
	router.GET("/ws", func(c *gin.Context) {
		handlers.WebSocketHandler(hub, c)
	})
}

