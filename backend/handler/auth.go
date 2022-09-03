package handler

import (
	"errors"
	"go-api/models"
	"go-api/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	user := &models.Users{}

	if err := c.BodyParser(user); err != nil {
		return c.JSON(utils.ErrorResponse(err, "JSON payload invalid"))
	}

	validated, err := loginFormValidation(user)
	if err != nil {
		return c.JSON(utils.ErrorResponse(err, "Login validation error"))
	}

	if !validated {
		return c.JSON(utils.ErrorResponse(err, "Login is not validated"))
	}

	auth, err := user.Authenticate(c.Context())
	if err != nil {
		return c.JSON(utils.ErrorResponse(err, "Authentication failed"))
	}

	if !auth {
		return c.JSON(utils.ErrorResponse(
			errors.New("User not found"),
			"Invalid credential",
		))
	}

	sess, err := h.Session.Get(c)
	if err != nil {
		return c.JSON(utils.ErrorResponse(
			errors.New("Session configuration invalid"),
			"Session configuration invalid",
		))
	}

	// NOTE: this is just an example
	// The best practice is using encrypted session data
	sess.Set("username", user.Username)

	if err := sess.Save(); err != nil {
		return c.JSON(utils.ErrorResponse(
			errors.New("Session configuration invalid"),
			"Unable to saved session",
		))
	}

	return c.JSON(map[string]interface{}{
		"login":   auth,
		"message": "Successfully login",
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	sess, err := h.Session.Get(c)
	if err != nil {
		panic(err)
	}

	if err := sess.Destroy(); err != nil {
		return c.JSON(utils.ErrorResponse(err, "The is user unable to logout"))
	}

	return c.JSON(map[string]bool{
		"logout": true,
	})
}

func (h *Handler) IsAuthenticated(c *fiber.Ctx) error {
	sess, err := h.Session.Get(c)
	if err != nil {
		panic(err)
	}

	// Check if user is authenticated
	login := sess.Get("username") != nil

	return c.JSON(map[string]bool{
		"login": login,
	})
}

func loginFormValidation(user *models.Users) (bool, error) {
	if len(user.Username) == 0 {
		return false, errors.New("Username is required")
	}

	if len(user.Password) == 0 {
		return false, errors.New("Password is required")
	}

	return true, nil
}
