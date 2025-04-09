package health

import (
	"net/http"

	"__APP__/internal/services"
)

func NewHandler(healthSvc services.HealthServicer) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleGet)

	return mux
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	status := healthSvc.Check()
	w.Write([]byte("Stay healthy!! " + status))
}
