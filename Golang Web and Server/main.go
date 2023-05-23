package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// server connection
	db, err := sql.Open("mysql", "root:2023@(127.0.0.1:3306)/rony?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// read data form server
	var (
		id        int
		postTitle string
		content   string
	)

	query := "SELECT id, title, content FROM post WHERE id = 1"
	if err := db.QueryRow(query).Scan(&id, &postTitle, &content); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter() //

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello and welcome to golang server\n use gorilla mux router services")
	})
	r.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, postTitle+"\n"+content)
	})

	http.ListenAndServe(":4001", r) // fist parameter port number
}
