package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

//go:embed templates
var templateFS embed.FS

const BASE_LAYOUT = "base"

type View struct {
	Template *template.Template
	Layout   string
}

var funcs = template.FuncMap{
	"GetYear": func() int {
		return time.Now().Year()
	},
}

func getLayoutFiles() []string {
	files, err := filepath.Glob("templates/*.layout.tmpl")
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(pages ...string) (*View, error) {
	files := getLayoutFiles()
	for _, f := range pages {
		files = append(files, fmt.Sprintf("templates/%s.page.tmpl", f))
	}
	var t *template.Template
	var err error

	_, exists := cache[pages[0]]

	if !exists {
		if env == "dev" {
			t, err = template.New("").Funcs(funcs).ParseFiles(files...)
		} else {
			t, err = template.New("").Funcs(funcs).ParseFS(templateFS, files...)
			cache[pages[0]] = t
		}
	} else {
		t = cache[pages[0]]
	}

	if err != nil {
		return nil, err
	}
	return &View{
		Template: t,
		Layout:   BASE_LAYOUT,
	}, nil
}

func (v *View) Render(w http.ResponseWriter, data any) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
