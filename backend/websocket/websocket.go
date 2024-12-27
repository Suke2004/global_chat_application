package customwebsocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin allows all origins by default.
	CheckOrigin: func(r *http.Request) bool {
		// You can add more robust origin checking here if needed.
		return true
	},
}

// Upgrade upgrades an HTTP connection to a WebSocket connection.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection to WebSocket: %v\n", err)
		return nil, fmt.Errorf("websocket upgrade error: %w", err)
	}
	log.Printf("WebSocket connection established from %s\n", r.RemoteAddr)
	return conn, nil
}

