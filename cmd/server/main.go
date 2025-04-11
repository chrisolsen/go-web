package main

import (
	"chrisolsen-goweb/internal/services"
	"chrisolsen-goweb/internal/transport/http"
	"fmt"
)

func main() {
	app := &http.App{
		AuthSvc:    services.NewAuthService(),
		HealthSvc:  services.NewHealthService(),
		EmailSvc:   services.NewEmailService(),
		PaymentSvc: services.NewPaymentService(),
		LoggingSvc: services.NewLoggingService(),
	}

	if err := app.Run(); err != nil {
		fmt.Printf("Failed to start application: %s", err)
	}
}
