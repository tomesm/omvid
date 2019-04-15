package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/tomesm/virtd/pkg/models"
)

type templateData struct {
	CurrentYear int
	Course      *models.Course
	Courses     []*models.Course
}

// Return nicely formatted string of time.Time object
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	// Get slice of all 'page' templates for the app
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fname := filepath.Base(page)
		println(page)
		//Register template.FuncMap
		ts, err := template.New(fname).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// Add all layout/master templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.master.tmpl"))
		if err != nil {
			return nil, err
		}
		// Add all partials
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[fname] = ts
	}
	return cache, nil
}
