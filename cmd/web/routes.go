package main

import (
	"net/http"
)

func (app *App) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	status := app.services.HealthSvc.Check()
	result := "Stay healthy!!! " + status
	w.Write([]byte(result))
}
