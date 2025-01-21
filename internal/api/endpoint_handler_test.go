package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestEndpointHandler tests the 4 different possibilities handled by the EndpointHandler
func TestEndpointHandler(t *testing.T) {
	numbers := []int{0, 10, 20, 30, 40, 50}

	tests := []struct {
		name               string
		target             string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "Valid target found",
			target:             "20",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"status":200,"index":2,"value":20}`,
		},
		{
			name:               "Valid target not found, within margin",
			target:             "25",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"status":200,"errorMessage":"Value not found within acceptable margin"}`,
		},
		{
			name:               "Invalid target",
			target:             "invalid",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"status":400,"errorMessage":"Invalid parameter. Must be an integer."}`,
		},
		{
			name:               "Valid target out of range",
			target:             "1000",
			expectedStatusCode: http.StatusOK,
			expectedBody:       `{"status":200,"errorMessage":"Value not found within acceptable margin"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/endpoint/"+tt.target, nil)
			rec := httptest.NewRecorder()

			handler := EndpointHandler(numbers)
			handler.ServeHTTP(rec, req)

			resp := rec.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tt.expectedStatusCode, resp.StatusCode)
			}

			body := strings.TrimSpace(rec.Body.String())
			if body != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, body)
			}
		})
	}
}
