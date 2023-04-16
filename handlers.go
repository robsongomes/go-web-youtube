package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type TemplateData struct {
	Email    string
	Telefone string
	Route    string
	Errors   []string
	User     *SessionUser
}

type SessionUser struct {
	Email string
}

func getUserFromCookie(r *http.Request) *SessionUser {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	return &SessionUser{Email: cookie.Value}
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
				Error    string `json:"error"`
			}

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Println(err)
			}

			user, err := FindUserByEmail(data.Email)
			if err != nil {
				log.Println(err)
			}

			if data.Password == user.Password {
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

func (app *Application) SignupHandler(view *View) http.HandlerFunc {
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
				Error    string `json:"error"`
			}

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Println(err)
			}

			//salvar o usuário no banco de dados

			stmt, err := db.Prepare("insert into users (email, password) values (?, ?)")
			if err != nil {
				data.Error = err.Error()
				json.NewEncoder(w).Encode(data)
				return
			}

			_, err = stmt.Exec(data.Email, data.Password)
			if err != nil {
				data.Error = err.Error()
				json.NewEncoder(w).Encode(data)
				return
			}

			//loga o usuário recém cadastrado
			data.Success = true
			data.Error = ""

			cookie := http.Cookie{
				Name:     "session",
				Value:    data.Email,
				Expires:  time.Now().Add(time.Minute * 3),
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)

			json.NewEncoder(w).Encode(data)
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

func (app *Application) NewPostHandler(view *View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := view.Render(w, r, nil)
			if err != nil {
				log.Println(err)
			}
		} else if r.Method == http.MethodPost {
			title := r.FormValue("title")
			content := r.FormValue("content")
			errors := make([]string, 0)

			if len(strings.Trim(title, " ")) == 0 {
				errors = append(errors, "Title is required")
			}

			if len(strings.Trim(content, " ")) == 0 {
				errors = append(errors, "Content is required")
			}

			userDTO := getUserFromCookie(r)
			user, err := FindUserByEmail(userDTO.Email)
			if err != nil {
				errors = append(errors, "You are not logged in")
			}

			post := Post{
				Title:   title,
				Content: content,
				Slug:    slugify(title),
				Author:  user,
			}

			log.Println(post)

			if len(errors) > 0 {
				view.Render(w, r, &TemplateData{Errors: errors})
				return
			}

			err = view.Render(w, r, nil)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func slugify(value string) string {
	value = strings.ToLower(value)
	reg := regexp.MustCompile("[^a-z0-9]+")
	value = reg.ReplaceAllString(value, "-")
	value = strings.Trim(value, "-")
	return value
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
