package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// NOTE: Fiber doesn't specifically define middleware types.
// All we need to todo is create a simple handler but define it
// on `app.Use(yourmiddleware)` method.

// AuthenticationMiddleware middleware for authentication
func AuthenticationMiddleware(c *fiber.Ctx) error {
	log.Println("auth_middleware")
	return c.Next()
}
