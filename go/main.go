package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting database connection...")
	db, err := connect_sqlite()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err) // Use Fatalf for errors
	}
	defer db.db.Close() // Close the connection in defer, after it is used
	hashPassword("test")
	hashPassword("1234")
	hashPassword("ji3g4go6?")
	fmt.Println("We now have a database connection and can use it")
	_, ok := db.insert_value_User("boss", "ej03xu35k3")
	if ok {
		fmt.Println("Sucessfully inserted user")
	}
	// Call show all users here
	allUsernames, ok := db.show_all_User()
	if ok {
		fmt.Println("printing usernames:")
		for _, userName := range allUsernames {
			log.Println(userName)
		}
	}

	// Serve Static Files
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("../CSS"))))
	http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("../JS"))))
	http.Handle("/IMAGE/", http.StripPrefix("/IMAGE/", http.FileServer(http.Dir("../IMAGE"))))
	http.Handle("/HTML/", http.StripPrefix("/HTML/", http.FileServer(http.Dir("../HTML"))))

	// Redirect to homepage
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/HTML/index.html", http.StatusFound)
	})

	// Start Server
	log.Println("Server is listening on: http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
