package errors

import (
	"errors"
	"fmt"

	"github.com/Alwin18/golang-module-template/internal/shared/constants"
	"github.com/Alwin18/golang-module-template/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

// AppError represents a structured application error with HTTP status code
type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Details    map[string]any
}

func (e *AppError) Error() string {
	return e.Message
}

// Error mapping untuk standard errors dari constants package
var errorStatusMap = map[error]int{
	constants.ErrDataNotFound:              fiber.StatusNotFound,
	constants.ErrAccountNotFound:           fiber.StatusNotFound,
	constants.ErrAccountInactive:           fiber.StatusNotFound,
	constants.ErrInvalidPassword:           fiber.StatusUnauthorized,
	constants.ErrInvalidRequest:            fiber.StatusBadRequest,
	constants.ErrNimAlreadyExists:          fiber.StatusConflict,
	constants.ErrEmailAlreadyExists:        fiber.StatusConflict,
	constants.ErrInternalServer:            fiber.StatusInternalServerError,
	constants.ErrAccountNotFoundOrInactive: fiber.StatusNotFound,
}

// Constructors for AppError
func NewNotFoundError(resource, identifier string) *AppError {
	return &AppError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s with %s not found", resource, identifier),
		StatusCode: fiber.StatusNotFound,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		StatusCode: fiber.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:       "UNAUTHORIZED",
		Message:    message,
		StatusCode: fiber.StatusUnauthorized,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    message,
		StatusCode: fiber.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:       "BAD_REQUEST",
		Message:    message,
		StatusCode: fiber.StatusBadRequest,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:       "FORBIDDEN",
		Message:    message,
		StatusCode: fiber.StatusForbidden,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:       "CONFLICT",
		Message:    message,
		StatusCode: fiber.StatusConflict,
	}
}

// GetStatusCode returns the appropriate HTTP status code for an error
// This function handles both AppError and standard errors from constants package
func GetStatusCode(err error) int {
	if appErr, ok := err.(*AppError); ok {
		return appErr.StatusCode
	}

	for knownErr, statusCode := range errorStatusMap {
		if errors.Is(err, knownErr) {
			return statusCode
		}
	}

	return fiber.StatusInternalServerError
}

// HandleError is a global error handler for Fiber handlers
// Usage in handler:
//
//	if err != nil {
//	    return errors.HandleError(ctx, err)
//	}
func HandleError(ctx *fiber.Ctx, err error) error {
	statusCode := GetStatusCode(err)

	return ctx.Status(statusCode).JSON(response.NewErrorResponse(response.ResponseError{
		Message: err.Error(),
		Code:    statusCode,
	}))
}

// HandleErrorWithDetails is similar to HandleError but allows adding custom error details
func HandleErrorWithDetails(ctx *fiber.Ctx, err error, details []string) error {
	statusCode := GetStatusCode(err)

	return ctx.Status(statusCode).JSON(response.NewErrorResponse(response.ResponseError{
		Message: err.Error(),
		Code:    statusCode,
		Errors:  details,
	}))
}

// RegisterErrorMapping allows dynamic registration of new error-to-status mappings
// This is useful for module-specific errors
func RegisterErrorMapping(err error, statusCode int) {
	errorStatusMap[err] = statusCode
}
