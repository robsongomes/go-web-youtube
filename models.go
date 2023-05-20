package main

import (
	"fmt"
	"html/template"
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
	Content   template.HTML
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
	rows, err := db.Query(`select p.id, p.title, p.slug, p.content, p.user_id, u.email, p.created_at, p.updated_at 
							from posts p join users u on p.user_id = u.id`)
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
			&user.Email,
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

func GetPostById(id int) (*Post, error) {
	var post Post
	var user User
	row := db.QueryRow(`select p.id, p.title, p.slug, p.content, p.user_id, u.email, p.created_at, p.updated_at 
	from posts p join users u on p.user_id = u.id where p.id = ?`, id)
	err := row.Scan(
		&post.Id,
		&post.Title,
		&post.Slug,
		&post.Content,
		&user.Id,
		&user.Email,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	post.Author = &user
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePost(post Post) error {
	stmt, err := db.Prepare("update posts set title = ?, content = ?, slug = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(post.Title, post.Content, post.Slug, post.UpdatedAt, post.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int) error {
	res, err := db.Exec("delete from posts where id = ?", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows != 1 {
		return fmt.Errorf("post %d does not exist", id)
	}
	return nil
}
