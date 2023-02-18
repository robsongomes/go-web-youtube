package main

import "net/http"

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "index")
}

func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "contact")
}

func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "about")
}

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "login")
}
