package main

// https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d

import (
	"chrisolsen-goweb/apps"
	"chrisolsen-goweb/internal/services"
	"log"
	"net/http"
	"text/template"
)

type App struct {
	apps.Base
	services Services
	state    State
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
	app := &App{
		services: Services{
			Auth:    services.NewAuthenticator(),
			Email:   services.NewEmailer(),
			Payment: services.NewPaymenter(),
			Log:     services.NewLogger(),
		},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("apps/public/layouts/base.html")
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	app.Run(":3000", mux)
}
