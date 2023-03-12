package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
	Errors   []string
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
		var email string
		cookie, _ := r.Cookie("session")
		if cookie != nil {
			email = cookie.Value
		}

		err := view.Render(w, TemplateData{Route: "about", Email: email})
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) LoginHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := view.Render(w, TemplateData{Route: "login"})
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {

			var data struct {
				Email    string `json:"email"`
				Password string `json:"password"`
				Success  bool   `json:"success"`
			}

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Println(err)
			}

			if data.Password == "123456" {
				//login com sucesso
				data.Success = true

				cookie := http.Cookie{
					Name:     "session",
					Value:    data.Email,
					Expires:  time.Now().Add(time.Minute * 3),
					HttpOnly: true,
				}

				http.SetCookie(w, &cookie)

				json.NewEncoder(w).Encode(data)
				return
			} else {
				json.NewEncoder(w).Encode(data)
				return
			}
		}
	}
}

// func (app *Application) LoginHandler(view *View) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			err := view.Render(w, TemplateData{Route: "login"})
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		} else if r.Method == http.MethodPost {
// 			email := r.FormValue("email")
// 			password := r.FormValue("password")

// 			fmt.Println(email)

// 			if password == "123456" {
// 				//login com sucesso
// 				//redirecionar o usuário para home.
// 				http.Redirect(w, r, "/", http.StatusSeeOther)
// 				return
// 			}

// 			err := view.Render(w, TemplateData{Route: "login", Errors: []string{"Invalid credentials"}})
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	}
// }
