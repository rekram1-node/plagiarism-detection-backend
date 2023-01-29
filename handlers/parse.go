package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func compare(docs []string) (*PlagiarismResponse, error) {

	return &PlagiarismResponse{
		Documents: []string{
			"This is a default response",
		},
	}, nil
}

func getDocumentsFromBody(w http.ResponseWriter, r *http.Request) ([]string, error) {
	var req PlagiarismRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return req.Documents, nil
}
