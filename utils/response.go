package utils

import (
	"encoding/json"
	"net/http"
)

// DTResponse struct
type DTResponse struct {
	Message string     `json:"message,omitempty"`
	Data    []struct{} `json:"data,omitempty"`
}

// Response function
func Response(response http.ResponseWriter, statusCode int, data DTResponse) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(data)
}
