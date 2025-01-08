package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import the driver
)

type DB struct {
	db       *sql.DB
	filepath string
}

// CHANGED: added error return
func connect_sqlite() (*DB, error) {
	file := "../sql/order_db.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}

	fmt.Println("Successfully connected to SQLite database!")
	return &DB{db: db, filepath: file}, nil
}

// (db_name *DB)類似於class的部份的method
func (db_name *DB) insert_value_User(user_name string, user_password string) (*DB, bool) {
	if db_name == nil || user_password == "" || user_name == "" {
		log.Println("Invalid input: db_name is nil or user_password/user_name is empty")
		log.Println("User name:", user_name, "User password:", user_password)
		return nil, false
	}

	// 准備插入語句
	insertStmt, err := db_name.db.Prepare("INSERT INTO User (user_name, user_password, user_salt) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error preparing insert statement:", err)
		return nil, false
	}
	defer insertStmt.Close()

	// 生成 Hash 和 Salt
	password, salt, err := hashPassword(user_password)
	if err != nil {
		log.Println("Error generating hash and salt:", err)
		return nil, false
	}
	log.Println("Generated Password:", password, "Generated Salt:", salt)

	// 執行插入
	_, err = insertStmt.Exec(user_name, password, salt)
	if err != nil {
		log.Println("Error executing insert statement:", err)
		return nil, false
	}

	return db_name, true
}

func (db *DB) show_User(username string) (string, bool) {
	if db == nil || username == "" {
		log.Println("error input db: ", db.filepath)
		log.Println("or error username: ", username)
		return "", false
	}
	// CHANGED: Removed .Open(), just use existing connection
	// db.db.Open("sqlite3", file) //Remove this as well

	searchStmt, err := db.db.Prepare("SELECT user_name FROM User WHERE user_name = ?")
	if err != nil {
		log.Println("Error preparing select statement", err)
		return "", false
	}
	defer func() {
		if err := searchStmt.Close(); err != nil {
			log.Println("Error closing prepared statement", err)
		}
	}()

	row := searchStmt.QueryRow(username)
	var userName string
	err = row.Scan(&userName)

	if err == sql.ErrNoRows {
		log.Println("User Not Found", err)
		return "", false
	}

	if err != nil {
		log.Println("Error executing select statement", err)
		return "", false
	}
	// CHANGED: DO NOT close the database here
	// db.db.Close()
	return userName, true
}
func (db *DB) show_all_User() ([]string, bool) {
	if db == nil {
		log.Println("db not exist")
		return nil, false
	}
	searchStmt, err := db.db.Prepare("SELECT user_name FROM User")
	if err != nil {
		log.Println("Error preparing select statement", err)
		return nil, false
	}
	defer func() {
		if err := searchStmt.Close(); err != nil {
			log.Println("Error closing prepared statement", err)
		}
	}()

	rows, err := searchStmt.Query()
	if err != nil {
		log.Println("Error executing select statement", err)
		return nil, false
	}
	defer rows.Close()

	var userNames []string
	for rows.Next() {
		var userName string
		err = rows.Scan(&userName)
		if err != nil {
			log.Println("Error scanning row", err)
			return nil, false
		}
		userNames = append(userNames, userName)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating rows", err)
		return nil, false
	}
	return userNames, true
}
func (db *DB) verify_User_password(user_name string, user_input_password string) bool {
	if db == nil {
		return fmt.Errorf("database connection is nil") == nil
	}
	if user_name == "" || user_input_password == "" {
		log.Println("empty input username: ", user_name, " or input password: ", user_input_password)
		return false
	}
	username, operation_sucessful := db.show_User(user_name)
	if (username == user_name) && operation_sucessful {
		query := "SELECT user_password, user_salt FROM User WHERE user_name = ?"

		row := db.db.QueryRow(query, username)
		var user_hash_password string
		var user_salt string

		// 提取結果
		err := row.Scan(&user_hash_password, &user_salt)

		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("user not found") == nil
			}
			return fmt.Errorf("error querying database: %v", err) == nil
		}
		operation_sucessful = verifyPassword(user_input_password, user_hash_password, user_salt)
		return operation_sucessful
	}
	return false
}
func (db *DB) submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // 只接受 POST 請求
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm() // 解析表單
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username") // 讀取 username 參數
	password := r.FormValue("password") // 讀取 password 參數

	log.Println("Received POST request with username:", username, "and password:", password)
	operation_sucessful := db.verify_User_password(username, password)
	if operation_sucessful {
		// fmt.Fprintf(w, "<h1> Hello %s </h1>", username)
		http.ServeFile(w, r, `..\HTML\manage_home_page.html`)
	}
	return
}
