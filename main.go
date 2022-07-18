package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	w.Write([]byte("Home route."))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)

		return
	}

	fmt.Fprintf(w, "Snippet number %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not allowed", 405)

		return
	}

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
