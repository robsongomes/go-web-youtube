package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"

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

func createViews() {
	var err error

	LoginView, err = NewView("login")
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
}

func main() {
	cache := make(map[string]*template.Template)

	cfg := Config{Version: "1.0.0"}
	flag.StringVar(&cfg.Port, "port", "3000", "porta do servidor")
	flag.StringVar(&cfg.Env, "env", "dev", "ambiente de execução")

	flag.Parse()

	app := Application{
		Config: cfg,
		Cache:  cache,
	}

	//conectar com o mysql (mariadb)

	db, _ = sql.Open("mysql", "root:secret@/mysql")

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Conectou com o banco de dados")

	createViews()

	log.Printf("Servidor na versão %s-%s escutando na porta %s\n",
		cfg.Version, cfg.Env, cfg.Port)

	log.Fatal(app.Start())
}
