package main

import (
	"net/http"
)

func (app *App) rootHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, 200, "landing.page.html")
}
