package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the driver
)

type DB struct {
	db *sql.DB
}

func connect_sqlite() *DB {
	const file string = "../sql/order_db.db"

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return nil // Return nil if connection fails
	}

	fmt.Println("Successfully connected to SQLite database!")

	// Insert a new record
	insertStmt, err := db.Prepare("INSERT INTO User (user_name, user_password)  VALUES (?,?)")
	if err != nil {
		log.Println("Error preparing insert statement", err)
		db.Close()
		return nil
	}
	_, err = insertStmt.Exec("Boss", "xl3103") // 老闆
	if err != nil {
		log.Println("Error executing insert statement", err)
		db.Close()
		return nil
	}

	return &DB{db: db}
}
