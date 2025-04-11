package handlers

import (
	"net/http"

	"chrisolsen-goweb/internal/handlers/health"
	"chrisolsen-goweb/internal/services"
)

func NewAppHandler() http.Handler {
	mux := http.NewServeMux()

	// handlers
	healthHander := health.NewHandler(services.NewHealthService())

	// handle
	mux.Handle("/health", healthHander)

	return mux
}
