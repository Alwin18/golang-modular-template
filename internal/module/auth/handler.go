package auth

import (
	apperrors "github.com/Alwin18/golang-module-template/internal/shared/errors"

	response "github.com/Alwin18/golang-module-template/internal/shared/responses"
	"github.com/gofiber/fiber/v2"
)

// Handler handles HTTP requests for the auth module.
type Handler struct {
	service *Service
}

// NewHandler creates a new auth Handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	resp, err := h.service.Login(req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return response.Success(c, "login successful", resp)
}

func handleServiceError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*apperrors.AppError); ok {
		return response.Error(c, appErr.Code, appErr.Message, nil)
	}
	return response.Error(c, fiber.StatusInternalServerError, "internal server error", nil)
}
