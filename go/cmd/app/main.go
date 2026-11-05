// Command app wires the dependencies together. This is the only place that does.
package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kudrutechai-debug/gated-template/go/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /invoice", handler.Invoice)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on :%s", port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server stopped: %v", err)
	}
}
