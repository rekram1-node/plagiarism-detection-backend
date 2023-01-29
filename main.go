package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rekram1-node/plagiarism-detection-backend/handlers"
	"github.com/rekram1-node/text-processor/text"
)

const (
	listenAddr   = ":3000"
	w2vModelName = "w2vModel.bin"
)

func main() {
	model := loadModel(w2vModelName)
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Route("/plagiarism", func(r chi.Router) {
		r.Get("/find", handlers.PlagiarismFinderHandler(model))
		r.Post("/compare", handlers.PlagiarismComparisonHandler(model))
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

func loadModel(filename string) *text.Word2Vec {
	// something
	if os.Getenv("DEBUG") != "" {
		return &text.Word2Vec{}
	}

	model, err := text.LoadModel(filename)

	if err != nil {
		log.Fatal(err)
	}

	return model
}
