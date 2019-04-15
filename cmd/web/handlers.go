package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomesm/virtd/pkg/models"
)

// Define a home handler function which writes a byte slice containing response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	c, err := app.courses.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	td := &templateData{Courses: c}
	app.render(w, r, "home.page.tmpl", td)
}

// Displays a particular course based on its ID
func (app *application) showCourse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	c, err := app.courses.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	td := &templateData{Course: c}
	app.render(w, r, "show.page.tmpl", td)
}

// Creates a new course
func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// Dummy data
	title := "Linear algebra"
	content := "Advanced linear algebra course for computer science"
	expires := "7"

	id, err := app.courses.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}
	// Redirect the user to the newly created course record
	http.Redirect(w, r, fmt.Sprintf("/course?id=%d", id), http.StatusSeeOther)
}
