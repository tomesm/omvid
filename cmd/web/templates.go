package main

import (
	"html/template"
	"path/filepath"

	"github.com/tomesm/virtd/pkg/models"
)

type templateData struct {
	Course  *models.Course
	Courses []*models.Course
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

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add all layout/master templates
		ts, err := ts.ParseGlob(filepath.Join(dir, "*.master.tmpl"))
		if err != nil {
			return nil, err
		}

		// Add all partials
		ts, err := ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
