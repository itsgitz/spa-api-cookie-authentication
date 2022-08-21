package handler

import (
	"go-api/models"
	"go-api/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users := &models.Users{}

	getUsers, err := users.FindUsers(c.Context())
	if err != nil {
		return c.JSON(utils.ErrorResponse(err, "Unable to get users data"))
	}

	return c.JSON(getUsers)
}
