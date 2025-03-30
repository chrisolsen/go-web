package health

import (
	"net/http"

	"github.com/chrisolsen/template/internal/services"
	"github.com/go-chi/chi/v5"
)

type HealthApiHandler struct {
	healthSvc services.HealthServicer
}

func AddHandler(r chi.Router, healthSvc services.HealthServicer) {
	h := HealthApiHandler{
		healthSvc: healthSvc,
	}

	sr := chi.NewRouter()
	sr.Get("/", h.getStatus)
	sr.Get("/{id}", h.getStatusWithId)
	r.Mount("/health", sr)
}

func (hh HealthApiHandler) getStatus(w http.ResponseWriter, r *http.Request) {
	status := hh.healthSvc.Check()
	w.Write([]byte(status))
}

func (hh HealthApiHandler) getStatusWithId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte("Super Healthy!!" + id))
}
