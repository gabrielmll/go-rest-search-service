package api

import (
	"encoding/json"
	"net/http"
)

// Response is the JSON structure returned by the endpoint
type Response struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

func EndpointHandler(numbers []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		index := 10
		value := numbers[index]

		response := Response{
			Index: index,
			Value: value,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
