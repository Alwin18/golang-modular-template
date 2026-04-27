package auth

import (
	"github.com/Alwin18/golang-module-template/internal/shared/errors"
	"github.com/Alwin18/golang-module-template/internal/shared/response"
	validation "github.com/Alwin18/golang-module-template/internal/shared/validations"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

// Handler handles HTTP requests for the auth module.
type Handler struct {
	service  *Service
	validate *validator.Validate
}

// NewHandler creates a new auth Handler.
func NewHandler(service *Service, validate *validator.Validate) *Handler {
	return &Handler{service: service, validate: validate}
}

func (h *Handler) Login(ctx *fiber.Ctx) error {
	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		}))
	}

	if err := h.validate.Struct(body); err != nil {
		validationErrors := validation.FormatValidationErrors(err, body)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
			Message: "body request tidak sesuai",
			Code:    fiber.StatusBadRequest,
			Errors:  validationErrors,
		}))
	}

	resp, err := h.service.Login(body)
	if err != nil {
		return errors.HandleError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewResponse(resp, "success login", fiber.StatusOK))
}
