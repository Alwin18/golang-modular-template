package errors

import (
	"errors"
	"net/http"
)

// AppError represents a structured application error.
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Sentinel errors
var (
	ErrUnauthorized = &AppError{Code: http.StatusUnauthorized, Message: "unauthorized"}
	ErrForbidden    = &AppError{Code: http.StatusForbidden, Message: "forbidden"}
	ErrNotFound     = &AppError{Code: http.StatusNotFound, Message: "resource not found"}
	ErrConflict     = &AppError{Code: http.StatusConflict, Message: "resource already exists"}
	ErrValidation   = &AppError{Code: http.StatusUnprocessableEntity, Message: "validation error"}
	ErrInternal     = &AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	ErrBadRequest   = &AppError{Code: http.StatusBadRequest, Message: "bad request"}
	ErrTooMany      = &AppError{Code: http.StatusTooManyRequests, Message: "too many requests"}
)

// New creates a new AppError.
func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Wrap wraps an underlying error with an AppError.
func Wrap(appErr *AppError, err error) *AppError {
	return &AppError{Code: appErr.Code, Message: appErr.Message, Err: err}
}

// IsNotFound checks if the error is a not-found error.
func IsNotFound(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == http.StatusNotFound
	}
	return false
}

// IsConflict checks if the error is a conflict error.
func IsConflict(err error) bool {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.Code == http.StatusConflict
	}
	return false
}
