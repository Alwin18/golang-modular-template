package user

import (
	"strconv"

	apperrors "github.com/Alwin18/golang-module-template/internal/shared/errors"
	response "github.com/Alwin18/golang-module-template/internal/shared/responses"

	"github.com/gofiber/fiber/v2"
)

// Handler handles HTTP requests for the user module.
type Handler struct {
	service *Service
}

// NewHandler creates a new user Handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	users, total, totalPage, err := h.service.List(page, perPage)
	if err != nil {
		return handleErr(c, err)
	}

	return response.Paginated(c, "users retrieved", users, response.Meta{
		Page:      page,
		PerPage:   perPage,
		Total:     total,
		TotalPage: totalPage,
	})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid user id", nil)
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		return handleErr(c, err)
	}

	return response.Success(c, "user retrieved", user)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid user id", nil)
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	user, err := h.service.Update(uint(id), req)
	if err != nil {
		return handleErr(c, err)
	}

	return response.Success(c, "user updated", user)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid user id", nil)
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return handleErr(c, err)
	}

	return response.NoContent(c)
}

func parseID(c *fiber.Ctx) (int, error) {
	return strconv.Atoi(c.Params("id"))
}

func handleErr(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*apperrors.AppError); ok {
		return response.Error(c, appErr.Code, appErr.Message, nil)
	}
	return response.Error(c, fiber.StatusInternalServerError, "internal server error", nil)
}
