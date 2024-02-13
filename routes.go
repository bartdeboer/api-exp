package main

import (
	"fmt"
	"log"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *log.Logger) {
	mux.Handle("/", serveStaticFiles())
	mux.HandleFunc("/submit-form", handleFormSubmission)
	mux.Handle("/hello", handleHello())
	mux.Handle("/schema/", http.StripPrefix("/schema", http.HandlerFunc(handleSchema)))
	mux.HandleFunc("/submit-pdf-form", handlePDFFormSubmission)
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
	fs := http.FileServer(http.Dir("static"))
	return http.StripPrefix("/", fs)
}
