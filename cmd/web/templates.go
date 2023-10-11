package main

import (
	"html/template"
	"path/filepath"
	"snippetbox.disconico/internal/models"
	"time"
)

type templateData struct {
	Snippet     models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
}

func humanDate(t time.Time) string {
	return t.Format("Mon 02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}
