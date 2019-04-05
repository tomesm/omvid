package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	var err error
	defer func() {
		if err != nil {
			app.serverError(w, err)
		}
	}()
	// Initialize template paths. Home must be the *first in the slice
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.master.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// Read a template file into a template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return
	}

	// Execute template set to write the template content as the response body. Dynamic data is nil
	// for now.
	err = ts.Execute(w, nil)
	if err != nil {
		return
	}
}

// Displays a particular course based on its ID
func (app *application) showCourse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific course with ID %d...", id)
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
