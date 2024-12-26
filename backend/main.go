package main

import (
	customwebsocket "chatapplication/websocket"
	"log"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool,w http.ResponseWriter, r *http.Request) {{
	log.Println("websocket is connected")
	conn, err := customwebsocket.Upgrade(w,r)
	if err != nil {
		log.Println(err)
		return
	}
	
	client := *customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes{
	log.Println("routes are established")
	pool := customwebsocket.NewPool()
	go pool.Start()

	http.HandleFunc("/api/ws", func(w http.ResponseWriter, r *http.Request){
		serverWs(pool,w,r)
	})
}

func main(){
	setupRoutes()
	http.ListenAndServe(":8080",nil)
}