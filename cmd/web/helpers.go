package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError writes an error message and stack trace for current goroutine to errorLog and sends
// a generic 500 Internal Server Error response
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and description
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is just a wrapper around clientError and sends a 404 Not Found reponse
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// render templates from the cache
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.seerverError(w, fmt.Errorf("The template %s does not exists", name))
		return
	}

	if err := ts.Execute(w, td); err != nil {
		app.seerverError(w, err)
	}

}
