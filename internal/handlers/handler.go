package handlers

import (
	"net/http"

	"__PKG__/internal/services"
	"__PKG__/internal/handlers/health"
)

func NewAppHandler() http.Handler {
	mux := http.NewServerMux()

	// handlers
	healthHander := health.NewHandler(services.NewHealthService())

	// handle
	mux.Handle("/health", healthHander)

	return mux
}
