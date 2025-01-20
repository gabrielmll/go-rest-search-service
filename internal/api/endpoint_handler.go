package api

import (
	"net/http"
	"strconv"
	"strings"

	"go-rest-search-service/internal/service"
)

func EndpointHandler(numbers []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.Path, "/")
		indexStr := urlParts[len(urlParts)-1]
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			SendResponse(
				w,
				http.StatusBadRequest,
				"Invalid parameter. Must be an integer.",
				nil, nil)
			return
		}

		value := service.GetNumber(numbers, index)

		SendResponse(
			w,
			http.StatusOK,
			"",
			&index,
			&value)
	}
}
