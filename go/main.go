package main

import (
	// "bufio"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "strings"
)

func main() {
	db := connect_sqlite()
	if db == nil {
		log.Println("Database connection failed, and will return.")
		return // if database connection failed, return from main
	}
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter username:")
	// username, _ := reader.ReadString('\n')
	// username = strings.TrimSpace(username)

	// fmt.Print("Enter Password:")
	// password, _ := reader.ReadString('\n')
	// password = strings.TrimSpace(password)

	// insert_value_User(db, username, password)

	defer db.db.Close()
	fmt.Println("We now have a database connection and can use it")

	// 提供靜態文件的路徑
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("../CSS"))))
	http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("../JS"))))
	http.Handle("/IMAGE/", http.StripPrefix("/IMAGE/", http.FileServer(http.Dir("../IMAGE"))))
	http.Handle("/HTML/", http.StripPrefix("/HTML/", http.FileServer(http.Dir("../HTML"))))

	// 將 "/" 路徑重定向到首頁
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/HTML/index.html", http.StatusFound)
	})

	// 啟動伺服器
	log.Println("伺服器正在執行：http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
