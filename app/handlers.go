package main

import (
	"net/http"
)

type rootData struct {
	Foo string
}

func (app *App) rootHandler(w http.ResponseWriter, r *http.Request) {

	data := rootData{}
	data.Foo = "FooBar"

	// FIXME: Views are being keyed by name, which could result in a conflict
	app.render(w, http.StatusOK, "landing.page.html", data)
}

func (app *App) partialHandler(w http.ResponseWriter, r *http.Request) {
	app.partial(w, http.StatusOK, "test.partial.html", nil)
}
