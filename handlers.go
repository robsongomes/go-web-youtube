package main

import (
	"log"
	"net/http"
)

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
}

func (app *Application) HomeHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, TemplateData{Route: "index"})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (app *Application) ContactHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, TemplateData{
			Email:    "teste@gmail.com",
			Telefone: "888888888",
			Route:    "contact",
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) AboutHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, TemplateData{Route: "about"})
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) LoginHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, TemplateData{Route: "login"})
		if err != nil {
			log.Println(err)
		}
	}
}
