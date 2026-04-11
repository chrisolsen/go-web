package main

import (
	"bytes"
	"log"
	"net/http"
)

func (app *App) render(w http.ResponseWriter, status int, page string, data any) {
	ts, ok := app.templateCache[page]
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error parsing template:", page)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *App) partial(w http.ResponseWriter, status int, page string, data any) {
	ts, ok := app.partialCache[page]
	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error parsing template:", page)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, data)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
