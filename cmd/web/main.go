// https://golangify.com/dynamic-html-templates

package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AndreyUA/go-snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// Get port number from environment
	addr := flag.String("addr", ":4200", "network address")
	// Get MySQL options from environment
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL database credentials")
	flag.Parse()

	// Create info logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create error logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// For loggin HTTP errors
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Use info logs
	infoLog.Printf("Server started on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
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
