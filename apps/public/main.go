package main

// https://dev.to/antonkuklin/golang-graceful-shutdown-3n6d

import (
	"chrisolsen-goweb/apps"
	"chrisolsen-goweb/internal/services"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
	"log"
	"net/http"
	"text/template"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

type App struct {
	apps.Base
	services Services
	state    State
}

type Services struct {
	Auth    services.Authenticator
	Email   services.Emailer
	Payment services.Paymenter
	Log     services.Logger
}

type State struct {
	foo string
}

func newThrottle() throttled.HTTPRateLimiterCtx {
	store, err := memstore.NewCtx(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 5,
	}
	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiterCtx{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	return httpRateLimiter
}

func main() {
	app := &App{
		services: Services{
			Auth:    services.NewAuthenticator(),
			Email:   services.NewEmailer(),
			Payment: services.NewPaymenter(),
			Log:     services.NewLogger(),
		},
	}

	th := newThrottle()
	mw := alice.New(
		th.RateLimit,
		nosurf.NewPure,
	)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("apps/public/layouts/base.html")
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})

	chain := mw.Then(mux)

	app.Run(":3000", chain)
}
