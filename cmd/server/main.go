package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/chrisolsen/template/internal/transport/http"
)

type App struct{}

func (a *App) Run() error {
	fmt.Println("App server running on http://localhost:3000")

	handler := transportHTTP.NewHandler()
	if err := http.ListenAndServe(":3000", handler); err != nil {
		fmt.Println("failed to setup handlers")
		return err
	}

	return nil
}

func main() {
	app := &App{}
	if err := app.Run(); err != nil {
		fmt.Printf("Failed to start application: %s", err)
	}
}
