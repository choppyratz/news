package models

import "fmt"

type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse -
func NewErrorResponse(message string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{Message: fmt.Sprintf(message, args...)}
}
