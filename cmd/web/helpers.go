package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Render templates
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		app.serverError(w, fmt.Errorf("Template %s is not exists", name))
		return
	}

	err := ts.Execute(w, td)

	if err != nil {
		app.serverError(w, err)
	}
}

// serverError: 500 server's error
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError: custom error for response
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound: 404 error
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
