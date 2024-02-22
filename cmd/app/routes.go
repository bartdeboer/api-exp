package main

import (
	"fmt"
	"net/http"

	htmlrenderer "github.com/bartdeboer/api-exp/internal/renderer/html"
	pdfrenderer "github.com/bartdeboer/api-exp/internal/renderer/pdf"
	"github.com/bartdeboer/api-exp/internal/renderer/renderer"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/", serveStaticFiles())
	mux.HandleFunc("/handle-form", HandleForm)
}

func serveStaticFiles() http.Handler {
	fs := http.FileServer(http.Dir("./static/"))
	return http.StripPrefix("/", fs)
}

func GetRenderer(renderer string) renderer.FormRenderer {
	switch renderer {
	case "html":
		return &htmlrenderer.Form{}
	case "pdf":
		return &pdfrenderer.Form{}
	}
	return nil
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	schemaFile := r.URL.Query().Get("schema")
	if schemaFile == "" {
		http.Error(w, "Schema parameter is missing", http.StatusBadRequest)
		return
	}

	renderer := r.URL.Query().Get("renderer")
	if renderer == "" {
		http.Error(w, "Renderer parameter is missing", http.StatusBadRequest)
		return
	}

	form := GetRenderer(renderer)
	if form == nil {
		http.Error(w, fmt.Sprintf("No such renderer: %s", renderer), http.StatusBadRequest)
		return
	}

	err := form.Load(schemaFile)
	if err != nil {
		http.Error(w, "Failed to load schema", http.StatusInternalServerError)
		return
	}

	form.Output(r, w)
}
