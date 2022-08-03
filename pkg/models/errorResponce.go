package models

import (
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	//TODO implement me
	panic("implement me")
}

// NewErrorResponse -
func NewErrorResponse(message string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{Message: fmt.Sprintf(message, args...)}
}

//func Error(w http.ResponseWriter, message string) {
//	w.WriteHeader(http.StatusInternalServerError)
//	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
//}
