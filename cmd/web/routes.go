package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/course", app.showCourse)
	mux.HandleFunc("/course/create", app.createCourse)
	// File Server for serving "./ui/static" directory and it's files.
	fs := http.FileServer(http.Dir("./ui/static/"))
	// Register file server to handle URL paths with "/static".
	// Strip the "/static prefix before the request reachest the file server
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	return app.logRequest(secureHeaders(mux))
}
