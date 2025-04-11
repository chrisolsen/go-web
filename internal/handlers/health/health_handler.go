package health

import (
	"net/http"

	"chrisolsen-goweb/internal/services"
	"chrisolsen-goweb/internal/transport/http/middleware"
)

type healthHandle struct {
	healthSvc services.HealthServicer
}

func NewHandler(healthSvc services.HealthServicer) *http.ServeMux {
	mux := http.NewServeMux()

	healthHandle := &healthHandle{
		healthSvc: healthSvc,
	}

	mw := middleware.New(middleware.ContentType("text/html"))

	mux.Handle("GET /", mw.Then(healthHandle.handleGet()))

	return mux
}

func (h *healthHandle) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		status := h.healthSvc.Check()
		result := "Stay healthy! " + status
		w.Write([]byte(result))
	}
}
