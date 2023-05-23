package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter() //

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello and welcome to golang server\n use gorilla mux router services")
	})

	http.ListenAndServe(":4001", r) // fist parameter port number
}
