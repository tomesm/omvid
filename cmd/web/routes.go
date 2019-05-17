package main

import (
	"net/http"

	"github.com/bmizerany/pat"  // REST router
	"github.com/justinas/alice" // Middleware chain lib
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/course/create", dynamicMiddleware.ThenFunc(app.createCourseForm))
	mux.Post("/course/create", dynamicMiddleware.ThenFunc(app.createCourse))
	mux.Get("/course/:id", dynamicMiddleware.ThenFunc(app.showCourse))

	// Static files do not need middleware
	fs := http.FileServer(http.Dir("./ui/static/"))
	// Register file server to handle URL paths with "/static".
	mux.Get("/static/", http.StripPrefix("/static", fs))
	return standardMiddleware.Then(mux)
}
