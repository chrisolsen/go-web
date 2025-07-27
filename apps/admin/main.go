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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the admin"))
	})

	app.Run(":3030", mux)
}
