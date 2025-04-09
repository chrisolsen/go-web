package handlers

import (
	"net/http"

	"__APP__/internal/services"
	"__APP__/internal/handlers/health"
)

func NewAppHandler() http.Handler {
	mux := http.NewServerMux()

	// handlers
	healthHander := health.NewHandler(services.NewHealthService())

	// handle
	mux.Handle("/health", healthHander)

	return mux
}
