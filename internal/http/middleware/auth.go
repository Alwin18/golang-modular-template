package middleware

import (
	"strings"

	"github.com/Alwin18/golang-module-template/internal/shared/response"
	"github.com/Alwin18/golang-module-template/internal/shared/security"
	"github.com/gofiber/fiber/v2"
)

// Auth returns a JWT validation middleware.
func Auth(jwtManager *security.JWTManager) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(response.ResponseError{
				Message: "authorization header required",
				Code:    fiber.StatusBadRequest,
			}))
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.NewErrorResponse(response.ResponseError{
				Message: "invalid authorization header format",
				Code:    fiber.StatusUnauthorized,
			}))
		}

		claims, err := jwtManager.ParseToken(parts[1])
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.NewErrorResponse(response.ResponseError{
				Message: "invalid or expired token",
				Code:    fiber.StatusUnauthorized,
			}))
		}

		// Store claims in context locals
		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("user_email", claims.Email)
		ctx.Locals("user_role", claims.Role)

		return ctx.Next()
	}
}

// RequireRole returns a middleware that enforces a specific role.
func RequireRole(role string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userRole, ok := ctx.Locals("user_role").(string)
		if !ok || userRole != role {
			return ctx.Status(fiber.StatusForbidden).JSON(response.NewErrorResponse(response.ResponseError{
				Message: "insufficient permissions",
				Code:    fiber.StatusForbidden,
			}))
		}
		return ctx.Next()
	}
}
