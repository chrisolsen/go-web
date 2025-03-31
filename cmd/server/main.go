package main

import (
	"fmt"
	"net/http"

	"__PKG__/internal/handlers"
)

type App struct{}

func (a *App) Run() error {
	handler := handlers.NewAppHandler()
	err := http.ListenAndServe(":__PORT__", handler); if err == nil {
		fmt.Println("App server running on http://localhost:__PORT__")
	}

	return err
}

func main() {
	app := &App{}
	if err := app.Run(); err != nil {
		fmt.Printf("Failed to start application: %s", err)
	}
}
