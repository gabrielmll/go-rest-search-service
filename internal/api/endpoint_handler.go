package api

import (
	"net/http"
	"strconv"
	"strings"

	"go-rest-search-service/internal/service"
)

// EndpointHandler handles HTTP requests to search for a target value in a slice of numbers.
func EndpointHandler(numbers []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.Path, "/")
		targetStr := urlParts[len(urlParts)-1]
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			SendResponse(
				w,
				http.StatusBadRequest,
				"Invalid parameter. Must be an integer.",
				nil, nil)
			return
		}

		indexFound, valueFound, notFoundMessage := service.BinarySearchValue(numbers, target)
		if notFoundMessage != "" {
			SendResponse(
				w,
				http.StatusOK,
				notFoundMessage,
				nil, nil)
			return
		}

		SendResponse(
			w,
			http.StatusOK,
			"",
			&indexFound,
			&valueFound)
	}
}
