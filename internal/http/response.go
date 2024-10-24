// Package httpresponse provides utilities for creating and sending HTTP responses
// in a JSON API format. It includes functions for sending successful responses
// with data and error responses with detailed error information.
package httpresponse

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func ResponseOk(w http.ResponseWriter, message string) {
	response := Response{
		Data: map[string]string{
			"url": message,
		},
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, message string, statusCode int) {
	errorResponse := ErrorResponse{
		Errors: []ErrorDetail{
			{
				Status: http.StatusText(statusCode),
				Detail: message,
			},
		},
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(errorResponse)
}
