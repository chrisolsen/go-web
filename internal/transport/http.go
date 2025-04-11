package transport

import (
	"fmt"
	"net/http"

	"chrisolsen-goweb/internal/handlers/health"
	"chrisolsen-goweb/internal/services"
)

type Http struct {
	AuthSvc    services.AuthServicer
	HealthSvc  services.HealthServicer
	EmailSvc   services.EmailServicer
	PaymentSvc services.PaymentServicer
	LoggingSvc services.LoggingServicer
}

func (a *Http) Run() error {
	router := a.NewRouter()
	err := http.ListenAndServe(":3000", router)
	if err == nil {
		fmt.Println("App server running on http://localhost:3000")
	}

	return err
}

func (a *Http) NewRouter() http.Handler {
	mux := http.NewServeMux()

	// handlers
	healthHander := health.NewHandler(a.HealthSvc)

	// handle
	mux.Handle("/health", healthHander)

	return mux
}
