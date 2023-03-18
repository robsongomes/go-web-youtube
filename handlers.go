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
	User     *User
}

type User struct {
	Email string
}

func getUserFromCookie(r *http.Request) *User {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	return &User{Email: cookie.Value}
}

func (app *Application) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromCookie(r)
		if user == nil {
			log.Println("User not logged in. Redirecting to login page...")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		log.Println("User logged in")
		next(w, r)
	}
}

func (app *Application) HomeHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (app *Application) ContactHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, r, &TemplateData{
			Email:    "teste@gmail.com",
			Telefone: "888888888",
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) AboutHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, r, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) LoginHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := view.Render(w, r, nil)
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

func (app *Application) PostHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := view.Render(w, r, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// func (app *Application) LoginHandler(view *View) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			err := view.Render(w, r, &TemplateData{Route: "login"})
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

// 			err := view.Render(w, r, TemplateData{Route: "login", Errors: []string{"Invalid credentials"}})
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	}
// }
