package errors

import (
	"fmt"

	"github.com/Alwin18/golang-modular-template/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Details    map[string]any
}

func (e *AppError) Error() string {
	return e.Message
}

// Constructors
func NewNotFoundError(resource, identifier string) *AppError {
	return &AppError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s with %s not found", resource, identifier),
		StatusCode: 404,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    message,
		StatusCode: 400,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:       "UNAUTHORIZED",
		Message:    message,
		StatusCode: 401,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    message,
		StatusCode: 500,
	}
}

// HTTP Error Handler
func HandleHTTPError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*AppError); ok {
		return c.Status(appErr.StatusCode).JSON(response.ResponseError{
			Message: appErr.Message,
			Code:    appErr.StatusCode,
		})
	}

	// Default error
	return c.Status(500).JSON(response.ResponseError{
		Message: "Internal server error",
		Code:    500,
	})
}
