package main

import (
	"chrisolsen-goweb/internal/services"
	"chrisolsen-goweb/internal/transport/http"
)

func main() {
	app := &http.App{
		AuthSvc:    services.NewAuthService(),
		HealthSvc:  services.NewHealthService(),
		EmailSvc:   services.NewEmailService(),
		PaymentSvc: services.NewPaymentService(),
		LoggingSvc: services.NewLoggingService(),
	}

	app.Run()
}
