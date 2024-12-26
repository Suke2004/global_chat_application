package customwebsocker

import (
	"fmt"
)

type Pool struct{
	Register chan *Client
	Unregister chan *Client
	clients map[*Client]bool
	Broadcast chan Message
}

func NewPool()*Pool{
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func(p *Pool) Start(){
	for {
		select{
		case client := <- pool.Register:
			p.Clients[client] = true
			fmt.Println("total connection pool:- ", len(p.Clients))
			for  k, _ := range p.Clients{
				fmt.Println(k)
				k.Conn.WriteJSON(Message{Type: 1, Body: "New  User Joined"})
			}
		case client := <- pool.Unregister:
			delete(p.Clients,client)
			fmt.Println("total connetion pool:- ",len(p.Clients))
			for k,_ := range p.Clients{
				fmt.Println(k)
				k.Conn.WriteJSON(Message{Type: 2, Body: "User Left"})
			}
		case msg := <- pool.Broadcast:
			fmt.Println("broadcasting a message")
			for k,_ := range p.Clients{
				if err := k.Conn.WriteJSON(msg); err != nil{
					fmt.Println(err)
					return
				}
			}
		}
		
	}

}