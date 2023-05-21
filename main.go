package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var env = "dev"
var cache = make(map[string]*template.Template)
var db *sql.DB

var LoginView *View
var AboutView *View
var ContactView *View
var HomeView *View
var PostView *View
var SignupView *View
var NewPostView *View
var EditPostView *View
var HomePostView *View

func createViews() {
	var err error

	LoginView, err = NewView("login")
	if err != nil {
		log.Fatal(err)
	}
	SignupView, err = NewView("signup")
	if err != nil {
		log.Fatal(err)
	}
	AboutView, err = NewView("about")
	if err != nil {
		log.Fatal(err)
	}
	ContactView, err = NewView("contact")
	if err != nil {
		log.Fatal(err)
	}
	HomeView, err = NewView("index")
	if err != nil {
		log.Fatal(err)
	}
	PostView, err = NewView("post")
	if err != nil {
		log.Fatal(err)
	}
	NewPostView, err = NewView("post-new")
	if err != nil {
		log.Fatal(err)
	}
	EditPostView, err = NewView("post-edit")
	if err != nil {
		log.Fatal(err)
	}
	HomePostView, err = NewView("post-view")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cache := make(map[string]*template.Template)

	cfg := Config{Version: "1.0.0"}
	flag.StringVar(&cfg.Port, "port", "3000", "porta do servidor")

	flag.Parse()

	app := Application{
		Config: cfg,
		Cache:  cache,
	}

	//conectar com o mysql (mariadb)
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pass, db_host, db_port, db_name)

	db, _ = sql.Open("mysql", dsn)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conectou com o banco de dados")

	initTables()

	createViews()

	log.Printf("Servidor na versão %s-%s escutando na porta %s\n",
		cfg.Version, env, cfg.Port)

	log.Fatal(app.Start())
}

func initTables() {
	log.Println("Criando as tabelas")
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id int not null auto_increment,
		email varchar(255) unique,
		password varchar(255),
		primary key (id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id int not null auto_increment,
		title varchar(255) not null,
		slug varchar(255) not null unique,
		content text,
		user_id int not null,
		created_at timestamp default current_timestamp(),
		updated_at timestamp default current_timestamp(),
		primary key (id),
		foreign key (user_id) references users(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
