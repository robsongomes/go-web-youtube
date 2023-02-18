package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (a *Application) RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFiles(
			"templates/"+page+".page.tmpl",
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

	t.Execute(w, nil)
}
