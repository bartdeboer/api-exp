package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bartdeboer/api-exp/internal/htmlrenderer"
	"github.com/bartdeboer/api-exp/internal/pdfrenderer"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/", serveStaticFiles())
	mux.HandleFunc("/submit-form", handleFormSubmission)
	mux.Handle("/hello", handleHello())
	mux.Handle("/schema/", http.StripPrefix("/schema", http.HandlerFunc(HandleSchemaForm)))
	mux.HandleFunc("/submit-pdf-form", HandlePDFFormSubmission)
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

func HandleSchemaForm(w http.ResponseWriter, r *http.Request) {
	schemaFile := r.URL.Path

	if !strings.HasSuffix(schemaFile, ".xml") {
		http.NotFound(w, r)
		return
	}

	var form htmlrenderer.Form
	err := form.Load(schemaFile)
	if err != nil {
		http.Error(w, "Failed to load schema", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	form.Output(r, w)
}

func HandlePDFFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	schemaFile := r.URL.Query().Get("schema")
	if schemaFile == "" {
		http.Error(w, "Schema parameter is missing", http.StatusBadRequest)
		return
	}

	var form pdfrenderer.Form
	err := form.Load(schemaFile)
	if err != nil {
		http.Error(w, "Failed to load schema", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"form_submission.pdf\"")
	form.Output(r, w)
}
