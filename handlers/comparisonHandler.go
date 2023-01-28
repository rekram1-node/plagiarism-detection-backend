package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
)

type PlagiarismComparison struct {
	Default string `json:"default"`
}

func PlagiarismComparisonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		w.Header().Set("Content-Type", "application/json")
		res := &PlagiarismFinder{
			Default: "I am a default response",
		}

		response, err := json.Marshal(res)
		if err != nil {
			logger.Info().Err(err).Msg("failed to marshal json response")
			writeResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}
