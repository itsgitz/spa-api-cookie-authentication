package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Login here",
	})
}
