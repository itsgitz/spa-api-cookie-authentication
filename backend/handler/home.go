package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Home(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"author":  "Anggit M Ginanjar",
		"message": "Hello from go-api project!",
	})
}
