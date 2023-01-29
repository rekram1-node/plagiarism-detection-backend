package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rekram1-node/text-processor/text"
)

func PlagiarismComparisonHandler(model *text.Word2Vec, debug bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		w.Header().Set("Content-Type", "application/json")

		if debug {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(mockResponse)
		}

		docs, err := getDocumentsFromBody(w, r)
		if err != nil {
			logger.Info().Err(err).Msg("failed to unmarshal request")
			writeResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		res, err := compare(docs, model)
		if err != nil {
			logger.Info().Err(err).Msg("failed to compare documents")
			writeResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
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
