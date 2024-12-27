package main

import (
	customwebsocket "chatapplication/websocket"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	address = ":8080"
)

func serverWs(pool *customwebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket connection established")
	conn, err := customwebsocket.Upgrade(w, r)
	if err != nil {
		log.Printf("Error during connection upgrade: %v", err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	go client.Read() // Use goroutine to handle reading from the client
}

func setupRoutes(pool *customwebsocket.Pool) {
	log.Println("Setting up routes")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWs(pool, w, r)
	})
}

func main() {
	pool := customwebsocket.NewPool()
	go pool.Start()
	setupRoutes(pool)

	// Graceful shutdown
	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		log.Println("Shutting down server...")
		if err := srv.Shutdown(nil); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	log.Printf("Starting server on %s", address)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
