package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/k1941996/student-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student api program"))
	})

	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	slog.Info("Server started at", slog.String("address", server.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server %s", err.Error())

		}
	}()

	<-done
	slog.Info("\nshutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown", slog.String("error", err.Error()))

	}

	slog.Info("Server gracefully shutdown")
	// connect to db
	// setup router
	// setup server
}
