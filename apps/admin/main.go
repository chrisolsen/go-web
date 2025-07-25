package main

import (
	"chrisolsen-goweb/apps"
	"chrisolsen-goweb/internal/services"
	"net/http"
)

type App struct {
	apps.Base
	services Services
}

type Services struct {
	Auth services.Authenticator
	Log  services.Logger
}

func main() {
	app := &App{
		services: Services{
			Auth: services.NewAuthenticator(),
			Log:  services.NewLogger(),
		},
	}

	mux := http.NewServeMux()

	app.Run(":8080", mux)
}
