package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-rest-search-service/internal/logger"
)

// TestLoggerMiddleware tests the LoggerMiddleware functionality
func TestLoggerMiddleware(t *testing.T) {
	var logOutput strings.Builder
	logger.SetOutput(&logOutput)

	// Dummy handler to use with the middleware
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	middleware := LoggerMiddleware(dummyHandler)

	// Fake a test request
	req := httptest.NewRequest("GET", "/test-endpoint", nil)
	rr := httptest.NewRecorder()
	middleware.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	// Verify that the logger captured the expected log
	expectedLog := "Executing endpoint - Method: GET, Path: /test-endpoint"
	if !strings.Contains(logOutput.String(), expectedLog) {
		t.Errorf("expected log message to contain %q, got %q", expectedLog, logOutput.String())
	}
}
