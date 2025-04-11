package main

import (
	"fmt"
	"net/http"

	"chrisolsen-goweb/internal/handlers"
)

type App struct{}

func (a *App) Run() error {
	handler := handlers.NewAppHandler()
	err := http.ListenAndServe(":3000", handler); if err == nil {
		fmt.Println("App server running on http://localhost:3000")
	}

	return err
}

func main() {
	app := &App{}
	if err := app.Run(); err != nil {
		fmt.Printf("Failed to start application: %s", err)
	}
}
