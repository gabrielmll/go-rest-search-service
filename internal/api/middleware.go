package api

import (
	"net/http"

	"go-rest-search-service/internal/logger"
)

// LoggerMiddleware logs the details of each incoming request
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Executing endpoint - Method: %s, Path: %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
