package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
)

func PlagiarismComparisonHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		w.Header().Set("Content-Type", "application/json")
		res := &PlagiarismResponse{
			Documents: []string{
				"This is a default response",
			},
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
