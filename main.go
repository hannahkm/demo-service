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
		err := getError()
		jsonResponse(w, http.StatusServiceUnavailable, err)
	})

	var handler http.Handler = mux

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
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

	if err := srv.Shutdown(ctx); err != nil {
		wrappedErr := fmt.Errorf("server forced to shutdown: %w", err)
		fmt.Println("Server forced to shutdown: error", wrappedErr)
	}

	fmt.Println("Server exited properly")
	os.Exit(0)

}

func getError() error {
	slog.Error("Random error in home handler", "error", "random service unavailable")
	err := fmt.Errorf("%d: %s", http.StatusServiceUnavailable, "Service temporarily unavailable")
	return err
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to encode JSON response", "error", err)
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
	}
}
