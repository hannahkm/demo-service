package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusServiceUnavailable, getError(r.Context()))
	})

	var handler http.Handler = mux

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		slog.Info("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			wrappedErr := fmt.Errorf("server failed to start: %w", err)
			fmt.Println("Server failed to start: error", wrappedErr)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	slog.Info("Shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		wrappedErr := fmt.Errorf("server forced to shutdown: %w", err)
		fmt.Println("Server forced to shutdown: error", wrappedErr)
	}

	fmt.Println("Server exited properly")
	os.Exit(0)

}

func getError(ctx context.Context) error {
	if ctx.Err() != nil {
		slog.Error("Context error", "error", ctx.Err())
		return ctx.Err()
	}
	slog.Error("Random error in home handler", "error", "random service unavailable")
	return fmt.Errorf("%d: %s", http.StatusServiceUnavailable, "Service temporarily unavailable")
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to encode JSON response", "error", err)
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
	}
}
