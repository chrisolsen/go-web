package main

import (
	"chrisolsen-goweb/internal/middleware"
	"net/http"

	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
)

func (app App) NewRouter() http.Handler {

	// middleware

	th := middleware.NewThrottle()
	mw := alice.New(
		th.RateLimit,
		nosurf.NewPure,
	)

	mux := http.NewServeMux()

	// static files
	fs := http.FileServer(http.Dir("./app/static"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// routes
	mux.HandleFunc("/", app.rootHandler)
	mux.HandleFunc("/signup", app.signupHandler)
	mux.HandleFunc("/login", app.loginHandler)
	mux.HandleFunc("/partial", app.partialHandler)

	return mw.Then(mux)
}
