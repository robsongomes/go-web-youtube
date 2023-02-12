package main

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

func RenderTemplate(w http.ResponseWriter, page string) {
	tp, err := template.ParseFiles("templates/" + page + ".page.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
	tp.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "about")
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":3000", nil)
}
