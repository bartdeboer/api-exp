package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Run(ctx context.Context, mux *http.ServeMux, args []string, stdout, stderr io.Writer) error {
	logger := log.New(stderr, "http: ", log.LstdFlags)

	// mux = http.NewServeMux()
	// addRoutes(mux, logger)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(logger, mux),
	}

	go func() {
		logger.Printf("Server listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "Error listening and serving: %s\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(stderr, "Error shutting down HTTP server: %s\n", err)
	}
	return nil
}
