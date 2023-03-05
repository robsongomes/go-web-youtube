package main

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	//poderia levar a criação dos handlers para o main
	mux.HandleFunc("/", app.HomeHandler(HomeView))
	mux.HandleFunc("/contact", app.ContactHandler(ContactView))
	mux.HandleFunc("/about", app.AboutHandler(AboutView))

	mux.HandleFunc("/login", app.LoginHandler(LoginView))

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	return mux
}
