package middleware

import "github.com/gofiber/fiber/v2"

// Security sets recommended HTTP security headers.
func Security() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-Frame-Options", "DENY")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		if c.Secure() {
			c.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		}
		return c.Next()
	}
}
