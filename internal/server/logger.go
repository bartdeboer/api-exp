package server

import (
	"log"
	"net/http"
)

func loggingMiddleware(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		}()
		h.ServeHTTP(w, r)
	})
}
