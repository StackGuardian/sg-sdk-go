package api

import (
	"fmt"
	"io"
	"net/http"
)

// ErrorType represents the type of error.
type ErrorType string

const (
	// ErrorTypeAPI indicates an error from the API.
	ErrorTypeAPI ErrorType = "api_error"
	// ErrorTypeNetwork indicates a network error.
	ErrorTypeNetwork ErrorType = "network_error"
	// ErrorTypeValidation indicates a validation error.
	ErrorTypeValidation ErrorType = "validation_error"
	// ErrorTypeConfiguration indicates a configuration error.
	ErrorTypeConfiguration ErrorType = "configuration_error"
	// ErrorTypeUnknown indicates an unknown error.
	ErrorTypeUnknown ErrorType = "unknown_error"
)

// Error represents an error from the StackGuardian SDK.
type Error struct {
	// Type is the type of error.
	Type ErrorType

	// Message is the error message.
	Message string

	// StatusCode is the HTTP status code (for API errors).
	StatusCode int

	// RequestID is the request ID from the API (if available).
	RequestID string

	// Err is the underlying error (if any).
	Err error
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("[%s] %d: %s", e.Type, e.StatusCode, e.Message)
	}
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

// Unwrap returns the underlying error.
func (e *Error) Unwrap() error {
	return e.Err
}

// NewAPIError creates a new API error from an HTTP response.
func NewAPIError(resp *http.Response) *Error {
	body, _ := io.ReadAll(resp.Body)
	return &Error{
		Type:       ErrorTypeAPI,
		Message:    string(body),
		StatusCode: resp.StatusCode,
		RequestID:  resp.Header.Get("X-Request-ID"),
	}
}

// IsNotFoundError checks if an error is a 404 Not Found error.
func IsNotFoundError(err error) bool {
	if apiErr, ok := err.(*Error); ok {
		return apiErr.StatusCode == http.StatusNotFound
	}
	return false
}

// IsUnauthorizedError checks if an error is a 401 Unauthorized error.
func IsUnauthorizedError(err error) bool {
	if apiErr, ok := err.(*Error); ok {
		return apiErr.StatusCode == http.StatusUnauthorized
	}
	return false
}
