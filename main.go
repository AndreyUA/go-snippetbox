package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	w.Write([]byte("Home route."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display snippet."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet."))
}

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
