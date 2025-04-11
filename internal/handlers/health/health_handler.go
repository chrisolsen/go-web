package health

import (
	"net/http"

	"chrisolsen-goweb/internal/services"
)

var svc services.HealthServicer

func NewHandler(healthSvc services.HealthServicer) *http.ServeMux {
	mux := http.NewServeMux()

	svc = healthSvc

	mux.HandleFunc("GET /", handle(handleGet))

	return mux
}


func handle(fn func() string ) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		result := fn()
		w.Write([]byte(result))
	}
}

func handleGet() string {
	status := svc.Check()
	return "Stay healthy! " + status
}
