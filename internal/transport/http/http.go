package http

import (
	"fmt"
	"net/http"

	"chrisolsen-goweb/internal/handlers/health"
	"chrisolsen-goweb/internal/services"
	"chrisolsen-goweb/internal/transport/http/middleware"
)

type App struct {
	AuthSvc    services.AuthServicer
	HealthSvc  services.HealthServicer
	EmailSvc   services.EmailServicer
	PaymentSvc services.PaymentServicer
	LoggingSvc services.LoggingServicer
}

func (a *App) Run() error {
	router := a.NewRouter()
	err := http.ListenAndServe(":3000", router)
	if err == nil {
		fmt.Println("App server running on http://localhost:3000")
	}

	return err
}

func (a *App) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mw := middleware.New(middleware.LoggingMiddleware)

	// handlers
	healthHander := health.NewHandler(a.HealthSvc)

	// handle
	mux.Handle("/health", mw.Then(healthHander))

	return mux
}
