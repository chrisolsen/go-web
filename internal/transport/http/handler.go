package http

import (
	"net/http"

	"github.com/chrisolsen/template/internal/services"
	"github.com/chrisolsen/template/internal/transport/http/health"
	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	rootRouter := chi.NewRouter()
	health.AddHandler(rootRouter, services.NewHealthService())

	return rootRouter
}
