package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomesm/virtd/pkg/forms"
	"github.com/tomesm/virtd/pkg/models"
)

// Define a home handler function which writes a byte slice containing response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	c, err := app.courses.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	td := &templateData{Courses: c}
	app.render(w, r, "home.page.tmpl", td)
}

// Displays a particular course based on its ID
func (app *application) showCourse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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
	td := &templateData{
		Course: c,
	}
	app.render(w, r, "show.page.tmpl", td)
}

// Handler for creating a form
func (app *application) createCourseForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// Creates a new course
func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := app.courses.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Course successfully created!")
	// Redirect the user to the newly created course record
	http.Redirect(w, r, fmt.Sprintf("/course/%d", id), http.StatusSeeOther)
}
