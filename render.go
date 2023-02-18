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

func (a *Application) RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFS(
			TemplateFS,
			"templates/"+page+".page.tmpl",
			"templates/navbar.layout.tmpl",
			"templates/base.layout.tmpl",
		)
		if err != nil {
			log.Println(err)
			return
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache hit")
		t = a.Cache[page]
	}

	contact := struct {
		Email    string
		Telefone string
	}{
		Email:    "robson@gmail.com",
		Telefone: "88 988888888",
	}

	t.Execute(w, contact)
}
