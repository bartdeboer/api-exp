package main

import (
	"fmt"
	"net/http"

	"github.com/bartdeboer/api-exp/internal/htmlrenderer"
	"github.com/bartdeboer/api-exp/internal/pdfrenderer"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/", serveStaticFiles())
	mux.HandleFunc("/submit-form", handleFormSubmission)
	mux.Handle("/hello", handleHello())
	mux.Handle("/schema/", http.StripPrefix("/schema", http.HandlerFunc(htmlrenderer.HandleSchemaForm)))
	mux.HandleFunc("/submit-pdf-form", pdfrenderer.HandlePDFFormSubmission)
}

func handleHello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Please send a valid form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Received form submission with name: %s", name)
}

func serveStaticFiles() http.Handler {
	fs := http.FileServer(http.Dir("./static/"))
	return http.StripPrefix("/", fs)
}
