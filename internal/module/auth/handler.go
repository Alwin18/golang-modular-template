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

func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	tokens, err := h.service.Register(req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return response.Created(c, "registered successfully", tokens)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return response.Success(c, "login successful", tokens)
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	tokens, err := h.service.Refresh(req)
	if err != nil {
		return handleServiceError(c, err)
	}

	return response.Success(c, "token refreshed", tokens)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req RefreshRequest
	_ = c.BodyParser(&req)

	_ = h.service.Logout(userID, req.RefreshToken)
	return response.NoContent(c)
}

func (h *Handler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	me, err := h.service.Me(userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	return response.Success(c, "user profile", me)
}

func handleServiceError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*apperrors.AppError); ok {
		return response.Error(c, appErr.Code, appErr.Message, nil)
	}
	return response.Error(c, fiber.StatusInternalServerError, "internal server error", nil)
}
