package api

import (
	"encoding/json"
	"net/http"
)

// Response is the JSON structure returned by the endpoint
type Response struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Index        *int   `json:"index,omitempty"`
	Value        *int   `json:"value,omitempty"`
}

func SendResponse(w http.ResponseWriter, statusCode int, errorMessage string, index *int, value *int) {
	response := Response{
		Status:       statusCode,
		ErrorMessage: errorMessage,
		Index:        index,
		Value:        value,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	return
}
