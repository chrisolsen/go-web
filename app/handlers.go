package main

import (
	"net/http"
)

type rootData struct {
	Foo string
}

func (app *App) rootHandler(w http.ResponseWriter, r *http.Request) {
	// prevent wildcard matches
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.renderOk(w, "landing/landing.page.html", nil)
}

func (app *App) partialHandler(w http.ResponseWriter, r *http.Request) {
	app.partial(w, http.StatusOK, "test.partial.html", nil)
}

func (app *App) signupHandler(w http.ResponseWriter, r *http.Request) {
	app.renderOk(w, "app/signup/signup.page.html", nil)
}

func (app *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	app.renderOk(w, "app/login/login.page.html", nil)
}
