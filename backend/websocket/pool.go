package customwebsocket

import (
	"fmt"
	"log"
	"sync"
)

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   chan Message
	Database    *Database
	clientMutex sync.RWMutex
}

func NewPool(db *Database) *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Database:   db,
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.clientMutex.Lock()
			p.Clients[client] = true
			p.clientMutex.Unlock()

			fmt.Println("User connected. Total pool size:", len(p.Clients))

			// Send previous messages to the new client
			messages, err := p.Database.GetMessages(10)
			if err == nil {
				for _, msg := range messages {
					client.Conn.WriteJSON(msg)
				}
			}

		case client := <-p.Unregister:
			p.clientMutex.Lock()
			delete(p.Clients, client)
			p.clientMutex.Unlock()
			fmt.Println("User disconnected. Total pool size:", len(p.Clients))

		case msg := <-p.Broadcast:
			log.Println("Broadcasting message:", msg)
			p.Database.SaveMessage("anonymous", msg.Body)

			p.clientMutex.RLock()
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(msg); err != nil {
					fmt.Println("Error sending message:", err)
				}
			}
			p.clientMutex.RUnlock()
		}
	}
}
