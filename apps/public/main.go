package main

// https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d

import (
	"chrisolsen-goweb/apps"
	"chrisolsen-goweb/internal/services"
	"net/http"
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
		w.Write([]byte("Hello"))
	})

	app.Run(":3000", mux)
}
