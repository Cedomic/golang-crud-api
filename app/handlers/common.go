package handlers

import (
	"encoding/json"
	"net/http"
)

// respondJSON makes the response with payload in JSON format
func respondJSON(responseWriter http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(err.Error()))
		return
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	responseWriter.Write([]byte(response))
}

// respondError makes the error response with payload in JSON format
func respondError(responseWriter http.ResponseWriter, code int, message string) {
	respondJSON(responseWriter, code, map[string]string{"error": message})
}
