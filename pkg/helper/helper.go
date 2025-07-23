package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Message    any `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{StatusCode: statusCode, Message: err.Error()}
}

func Unauthorized() APIError {
	return APIError{StatusCode: http.StatusUnauthorized, Message: "Unauthorized"}
}

func BadRequest() APIError {
	return APIError{StatusCode: http.StatusBadRequest, Message: "Bad Request"}
}

func InvalidJSON() APIError {
	return APIError{StatusCode: http.StatusBadRequest, Message: "Invalid JSON request body"}
}

func NotFound() APIError {
	return APIError{StatusCode: http.StatusNotFound, Message: "Not Found"}
}

func InternalServerError() APIError {
	return APIError{StatusCode: http.StatusInternalServerError, Message: "Internal Server Error"}
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"status_code":500,"message":"error encoding JSON"}`, http.StatusInternalServerError)
		return err
	}
	return nil
}

type ValidationError struct {
	Errors map[string]string
}

func (v ValidationError) Error() string {
	return "validation failed"
}
