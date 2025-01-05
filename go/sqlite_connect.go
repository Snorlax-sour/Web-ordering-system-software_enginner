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
	db.Close()
	return &DB{db: db}
}
func insert_value_User(db_name *DB, user_name string, user_password string) (*DB, bool /*operation ture or false*/) {
	if db_name == nil {
		log.Println("Error connecting to database:", db_name)
		return nil, false
	}
	// Insert a new record
	insertStmt, err := db_name.db.Prepare("INSERT INTO User (user_name, user_password, user_salt)  VALUES (?,?,?)")
	if err != nil {
		log.Println("Error preparing insert statement", err)
		return nil, false
	}
	password, salt, err := hashPassword(user_password)
	if err != nil {
		log.Println("generated hash password encounter error")
		return nil, false
	}

	_, err = insertStmt.Exec(user_name, password, salt) // insert into values
	if err != nil {
		log.Println("Error executing insert statement", err)
		return nil, false
	}

	return db_name, true

}
