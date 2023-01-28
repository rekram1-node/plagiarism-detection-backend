package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rekram1-node/plagiarism-detection-backend/handlers"
)

const (
	listenAddr = ":3000"
)

func main() {
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Route("/plagiarism", func(r chi.Router) {
		r.Get("/find", handlers.PlagiarismFinderHandler())
		r.Post("/compare", handlers.PlagiarismComparisonHandler())
	})

	server := &http.Server{
		Addr:              listenAddr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	logger.Info().Msg("Server Started")

	<-done
	logger.Info().Msg("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Panic().Err(err).Msg("server shutdown failed")
		return
	}

	logger.Info().Msg("Server Exited Properly")
}
