package main

import (
	"net/http"

	"github.com/bmizerany/pat"  // REST router
	"github.com/justinas/alice" // Middleware chain lib
)

func (app *application) routes() http.Handler {
	middleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/course/create", http.HandlerFunc(app.createCourseForm))
	mux.Post("/course/create", http.HandlerFunc(app.createCourse))
	mux.Get("/course/:id", http.HandlerFunc(app.showCourse))

	fs := http.FileServer(http.Dir("./ui/static/"))
	// Register file server to handle URL paths with "/static".
	mux.Get("/static/", http.StripPrefix("/static", fs))
	return middleware.Then(mux)
}
