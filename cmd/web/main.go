package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server started on http://localhost:4200")
	err := http.ListenAndServe(":4200", mux)
	log.Fatal(err)
}
