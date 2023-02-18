package main

import (
	"flag"
	"html/template"
	"log"
)

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

	log.Printf("Servidor na versão %s-%s escutando na porta %s\n",
		cfg.Version, cfg.Env, cfg.Port)

	log.Fatal(app.Start())
}
