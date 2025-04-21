package http

// https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func (app *App) Run() {
	router := app.NewRouter()
	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %s\n", err)
		}
	}()
	log.Println("Server started on :3000")

	// Wait for shutdown signal
	<-done
	log.Println("Server shutting down...")

	// Run shutdown logic
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %s\n", err)
		}
		cancel()
		log.Println("Server shutdown complete")
	}()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %s\n", err)
	}
}

// Shutdown handles the shutdown logic for the application
func (a *App) Shutdown() error {
	// TODO: implement shutdown logic
	// This could include closing database connections, stopping background jobs, etc.
	// For now, we'll just log that the shutdown is happening
	log.Println("Shutting down services...")
	return nil
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
