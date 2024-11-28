package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/internal/models"
)

func SendApiResponse(w http.ResponseWriter, response models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	// Encode the response as JSON and send it back to the client

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
