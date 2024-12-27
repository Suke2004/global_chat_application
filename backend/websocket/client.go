package customwebsocket

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type      int       `json:"type"`
	Body      string    `json:"body"`
	Sender    string    `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
	SenderID  string    `json:"sender_id"` // Optional: To identify sender
}

func (c *Client) Read() {
	defer func() {
		// Unregister client from pool and close the connection
		c.Pool.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()

	for {
		// Read incoming WebSocket message
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			// Log the error and break out of the loop
			log.Printf("Error reading message: %v", err)
			return
		}

		// Create message object to send through the broadcast channel
		m := Message{
			Type: msgType,
			Body: string(msg),
		}

		// Broadcast message to all clients in the pool
		c.Pool.Broadcast <- m
		fmt.Printf("Message received: %s\n", m.Body)
	}
}
