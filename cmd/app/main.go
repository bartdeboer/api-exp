package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bartdeboer/api-exp/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Cwd: %s\n", wd)

	mux := http.NewServeMux()
	addRoutes(mux)

	if err := server.Run(ctx, mux, os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
