package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello and welcome to golang server")
	})

	http.ListenAndServe(":4001", nil) // fist parameter port number
}
