package api

import (
	"encoding/json"
	"net/http"
)


type Response [T any] struct {
	Status string  `json:"status"`
	Message string  `json:"message"`
	Data T   `json:"data"`
}

func Success[T any] (w http.ResponseWriter, data T, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := Response[T] {
		Status: "success",
		Message: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(resp)
}

func ErrorResponse[T any](w http.ResponseWriter, message string, data ...T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	var payload T
	if len(data) > 0 {
		payload = data[0]
	}

	resp := Response[T]{
		Status:  "error",
		Message: message,
		Data:    payload,
	}

	json.NewEncoder(w).Encode(resp)
}

func BadResponse[T any](w http.ResponseWriter, message string, data ...T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	var payload T
	if len(data) > 0 {
		payload = data[0]
	}

	resp := Response[T]{
		Status:  "error",
		Message: message,
		Data:    payload,
	}

	json.NewEncoder(w).Encode(resp)
}