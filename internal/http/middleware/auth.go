package middleware

import (
	"strings"

	response "github.com/Alwin18/golang-module-template/internal/shared/responses"
	"github.com/Alwin18/golang-module-template/internal/shared/security"
	"github.com/gofiber/fiber/v2"
)

// Auth returns a JWT validation middleware.
func Auth(jwtManager *security.JWTManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, fiber.StatusUnauthorized, "authorization header required", nil)
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			return response.Error(c, fiber.StatusUnauthorized, "invalid authorization header format", nil)
		}

		claims, err := jwtManager.ParseToken(parts[1])
		if err != nil {
			return response.Error(c, fiber.StatusUnauthorized, "invalid or expired token", nil)
		}

		// Store claims in context locals
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// RequireRole returns a middleware that enforces a specific role.
func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("user_role").(string)
		if !ok || userRole != role {
			return response.Error(c, fiber.StatusForbidden, "insufficient permissions", nil)
		}
		return c.Next()
	}
}
