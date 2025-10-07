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

	// routes

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.rootHandler)

	return mw.Then(mux)
}
