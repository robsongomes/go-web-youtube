package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed templates
var TemplateFS embed.FS

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
}

var funcs = template.FuncMap{
	"GetYear": func() int {
		return time.Now().Year()
	},
}

func (a *Application) RenderTemplate(w http.ResponseWriter, page string, data any) error {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = parseTemplate(page, a.Config.Env)
		if err != nil {
			log.Println(err)
			return err
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache hit")
		t = a.Cache[page]
	}

	return t.ExecuteTemplate(w, "base", data)
}

func parseTemplate(page, env string) (*template.Template, error) {
	t := template.New("").Funcs(funcs)
	if env != "dev" {
		return t.ParseFS(
			TemplateFS,
			"templates/base.layout.tmpl",
			"templates/"+page+".page.tmpl",
			"templates/navbar.layout.tmpl",
		)
	}
	return t.ParseFiles(
		"templates/base.layout.tmpl",
		"templates/"+page+".page.tmpl",
		"templates/navbar.layout.tmpl",
	)
}
