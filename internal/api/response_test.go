package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendResponse(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		errorMessage   string
		index          *int
		value          *int
		expectedStatus int
		expectedBody   Response
	}{
		{
			name:           "Success Response",
			statusCode:     http.StatusOK,
			errorMessage:   "",
			index:          nil,
			value:          nil,
			expectedStatus: http.StatusOK,
			expectedBody: Response{
				Status:       http.StatusOK,
				ErrorMessage: "",
				Index:        nil,
				Value:        nil,
			},
		},
		{
			name:           "Error Response with message",
			statusCode:     http.StatusBadRequest,
			errorMessage:   "Invalid Request",
			index:          nil,
			value:          nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody: Response{
				Status:       http.StatusBadRequest,
				ErrorMessage: "Invalid Request",
				Index:        nil,
				Value:        nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock response
			rr := httptest.NewRecorder()

			SendResponse(rr, tt.statusCode, tt.errorMessage, tt.index, tt.value)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, status)
			}

			// Check the response body
			var actualBody Response
			if err := json.NewDecoder(rr.Body).Decode(&actualBody); err != nil {
				t.Fatalf("failed to decode response body: %v", err)
			}

			if actualBody != tt.expectedBody {
				t.Errorf("expected body %+v, got %+v", tt.expectedBody, actualBody)
			}
		})
	}
}
