package response

import "github.com/gofiber/fiber/v2"

// Meta holds pagination metadata.
type Meta struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

type successResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type paginatedResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
}

type errorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success sends a 200 JSON response.
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(successResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 JSON response.
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(successResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Paginated sends a paginated 200 JSON response.
func Paginated(c *fiber.Ctx, message string, data interface{}, meta Meta) error {
	return c.Status(fiber.StatusOK).JSON(paginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Error sends an error JSON response with the given status code.
func Error(c *fiber.Ctx, statusCode int, message string, errs interface{}) error {
	return c.Status(statusCode).JSON(errorResponse{
		Success: false,
		Message: message,
		Errors:  errs,
	})
}

// NoContent sends a 204 response.
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
