package main

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.HomeHandler)
	mux.HandleFunc("/contact", app.ContactHandler)
	mux.HandleFunc("/about", app.AboutHandler)

	mux.HandleFunc("/login", app.LoginHandler)

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	return mux
}
