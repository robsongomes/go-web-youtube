package main

import (
	"time"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type Post struct {
	Id        int
	Title     string
	Content   string
	Slug      string
	Author    *User
	CreatedAt time.Time
	UpdateAt  time.Time
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	//buscar o usuário pelo email
	row := db.QueryRow("select id, email, password from users where email = ?", email)

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return &user, err
	}
	return &user, nil
}
