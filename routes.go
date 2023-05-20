package main

import "net/http"

func (app *Application) Routes() http.Handler {
	mux := http.NewServeMux()

	//poderia levar a criação dos handlers para o main
	mux.HandleFunc("/", app.HomeHandler(HomeView))
	mux.HandleFunc("/contact", app.ContactHandler(ContactView))
	mux.HandleFunc("/about", app.AboutHandler(AboutView))
	mux.HandleFunc("/post", app.AuthMiddleware(app.PostHandler(PostView)))
	mux.HandleFunc("/post/view", app.HomePostViewHandler(HomePostView))
	mux.HandleFunc("/post/new", app.AuthMiddleware(app.NewPostHandler(NewPostView)))
	mux.HandleFunc("/post/edit", app.AuthMiddleware(app.EditPostHandler(EditPostView)))
	mux.HandleFunc("/post/delete", app.AuthMiddleware(app.DeletePostHandler))

	mux.HandleFunc("/login", app.LoginHandler(LoginView))
	mux.HandleFunc("/signup", app.SignupHandler(SignupView))
	mux.HandleFunc("/logout", app.LogoutHandler)

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	return mux
}
