package customwebsocket

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

// Initialize SQLite database
func NewDatabase(dbPath string) *Database {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to SQLite database: %v", err)
	}

	// Create necessary tables if they don't exist
	createMessagesTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender TEXT NOT NULL,
		message TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createConnectionsTable := `
	CREATE TABLE IF NOT EXISTS connections (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip_address TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute both table creation queries
	_, err = conn.Exec(createMessagesTable)
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	_, err = conn.Exec(createConnectionsTable)
	if err != nil {
		log.Fatalf("Failed to create connections table: %v", err)
	}

	return &Database{Conn: conn}
}

// Store message in the database
func (db *Database) SaveMessage(sender, message string) error {
	query := `INSERT INTO messages (sender, message) VALUES (?, ?);`
	_, err := db.Conn.Exec(query, sender, message)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
	}
	return err
}

// Fetch messages from the database
func (db *Database) GetMessages(limit int) ([]Message, error) {
	query := `SELECT sender, message, timestamp FROM messages ORDER BY id DESC LIMIT ?;`
	rows, err := db.Conn.Query(query, limit)
	if err != nil {
		log.Printf("Failed to fetch messages: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err = rows.Scan(&msg.Sender, &msg.Body, &msg.Timestamp)
		if err != nil {
			log.Printf("Failed to scan message: %v", err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

// Log connection metadata in the database
func (db *Database) LogConnection(ip string) error {
	query := "INSERT INTO connections (ip_address) VALUES (?)"
	_, err := db.Conn.Exec(query, ip)
	if err != nil {
		log.Printf("Failed to log connection: %v", err)
	}
	return err
}
