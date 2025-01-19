package api

import (
	"encoding/json"
	"net/http"
)

func EndpointHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("{'test': 'ok'}")
	}
}
