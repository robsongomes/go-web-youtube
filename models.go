package main

import (
	"log"
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
	UpdatedAt time.Time
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

func CreatePost(post Post) error {
	stmt, err := db.Prepare(`insert into posts (title, slug, content, user_id)
							values(?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(post.Title, post.Slug, post.Content, post.Author.Id)
	if err != nil {
		return err
	}
	return nil
}

func RetrievePosts() []Post {
	posts := []Post{}

	//buscar o usuário pelo email
	rows, err := db.Query(`select id, title, slug, content, user_id, created_at, updated_at from posts`)
	if err != nil {
		log.Println(err)
		return posts
	}

	for rows.Next() {
		var post Post
		var user User
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Slug,
			&post.Content,
			&user.Id,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return posts
		}
		post.Author = &user
		posts = append(posts, post)
	}

	return posts
}
