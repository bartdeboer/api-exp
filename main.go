package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func loggingMiddleware(logger *log.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		}()
		h.ServeHTTP(w, r)
	})
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) error {
	logger := log.New(stderr, "http: ", log.LstdFlags)

	mux := http.NewServeMux()
	addRoutes(mux, logger)

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

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx, os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
