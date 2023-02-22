package main

import "net/http"

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "index", nil)
}

func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "contact", TemplateData{
		Email:    "teste@gmail.com",
		Telefone: "888888888",
	})
}

func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "about", nil)
}

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "login", nil)
}
