package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Get port number from environment
	addr := flag.String("addr", ":4200", "network address")
	flag.Parse()

	// Create info logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create error logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	//Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Static files
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// For loggin HTTP errors
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Use info logs
	infoLog.Printf("Server started on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

// Open prevents to show static files
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)

	if err != nil {
		return nil, err
	}

	s, err := f.Stat()

	if s.IsDir() {
		index := filepath.Join(path, "index.html")

		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()

			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
