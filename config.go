package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Config struct {
	Port    string
	Env     string
	Version string
}

type Application struct {
	Config Config
	Cache  map[string]*template.Template
}

func (a *Application) Start() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", a.Config.Port),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           a.Routes(),
	}

	return srv.ListenAndServe()
}
