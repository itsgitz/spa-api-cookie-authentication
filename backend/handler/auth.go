package handler

import (
	"go-api/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	users := &models.Users{}

	if err := c.BodyParser(users); err != nil {
		return c.JSON(map[string]string{
			"error":   err.Error(),
			"message": "Error, something went wrong",
		})
	}

	auth, err := users.Authenticate(c.Context(), users.Username, users.Password)
	if err != nil {
		log.Println(err)
	}

	c.Cookie(setSession(users.Username))

	return c.JSON(map[string]interface{}{
		"login": auth,
	})
}

func setSession(username string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "session_id",
		Value:    username,
		SameSite: "None",
		HTTPOnly: true,
		Secure:   true,
	}
}
