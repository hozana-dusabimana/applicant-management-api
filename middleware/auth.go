package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
)

// SimpleAuth is a basic authentication middleware
func SimpleAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip auth for health check
		if c.Path() == "/health" {
			return c.Next()
		}

		// Get Authorization header
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authorization header required",
			})
		}

		// Check if it's a Bearer token
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		// Simple token validation (in real app, validate against database)
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == "" || len(token) < 10 {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Add user info to context (simplified)
		c.Locals("user_id", "user_123")
		c.Locals("user_role", "admin")

		return c.Next()
	}
}
