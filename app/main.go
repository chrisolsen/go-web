package main

// https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d

import (
	"chrisolsen-goweb/internal/services"
	"chrisolsen-goweb/internal/templates"
	"html/template"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type App struct {
	Base
	templateCache map[string]*template.Template
	partialCache  map[string]*template.Template
	services      Services
	state         State
}

type Services struct {
	Auth    services.Authenticator
	Email   services.Emailer
	Payment services.Paymenter
	Log     services.Logger
}

type State struct {
	foo string
}

func main() {
	templateCache, err := templates.NewTemplateCache("./app/views")
	if err != nil || len(templateCache) == 0 {
		log.Println("Failed to load template cache")
		return
	}

	partialCache, err := templates.NewPartialCache("./app/views")
	if err != nil || len(partialCache) == 0 {
		log.Println("Failed to load partial cache")
		return
	}

	app := &App{
		templateCache: templateCache,
		partialCache:  partialCache,
		services: Services{
			Auth:    services.NewAuthenticator(),
			Email:   services.NewEmailer(),
			Payment: services.NewPaymenter(),
			Log:     services.NewLogger(),
		},
	}

	router := app.NewRouter()

	app.Run(":3000", router)
}
