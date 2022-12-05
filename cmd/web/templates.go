package main

import (
	"github.com/server-practice/pkg/forms"
	"github.com/server-practice/pkg/models"
	"html/template"
	"path/filepath"
	"time"
)

type TemplateData struct {
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
	CurrentYear       int
	Form              *forms.Form
	Flash             string
	AuthenticatedUser *models.User
}

func humanDate(t time.Time) string {

	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan  2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplate(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil

}
