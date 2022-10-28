package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{Message: fmt.Sprintf(message, args...)}
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Message: message})
	if err != nil {
		fmt.Printf("failed encode: %v", err)
	}

}
