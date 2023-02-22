package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var TemplateFS embed.FS

type TemplateData struct {
	Email    string
	Telefone string
}

func (a *Application) RenderTemplate(w http.ResponseWriter, page string, data any) {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = parseTemplate(page, a.Config.Env)
		if err != nil {
			log.Println(err)
			return
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache hit")
		t = a.Cache[page]
	}

	t.ExecuteTemplate(w, "base", data)
}

func parseTemplate(page, env string) (*template.Template, error) {
	if env != "dev" {
		return template.ParseFS(
			TemplateFS,
			"templates/base.layout.tmpl",
			"templates/"+page+".page.tmpl",
			"templates/navbar.layout.tmpl",
		)
	}
	return template.ParseFiles(
		"templates/base.layout.tmpl",
		"templates/"+page+".page.tmpl",
		"templates/navbar.layout.tmpl",
	)
}
