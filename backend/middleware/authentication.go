package middleware

import (
	"errors"
	"go-api/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Middleware struct {
	Session *session.Store
}

func New(handlerSession *session.Store) *Middleware {
	m := &Middleware{}
	m.Session = handlerSession

	return m
}

// NOTE: Fiber doesn't specifically define middleware types.
// All we need to todo is create a simple handler but define it
// on `app.Use(yourmiddleware)` method.

// VerifyAuthentication middleware for authentication
func (m *Middleware) VerifyAuthentication(c *fiber.Ctx) error {
	userSession := getSession(c, m.Session)

	log.Println("User:", userSession, "is logged in")
	log.Println(c.Cookies("session_id"))

	if userSession == nil {
		return c.JSON(utils.ErrorResponse(
			errors.New("Invalid user session"),
			"The user is not authenticated",
		))
	}

	return c.Next()
}

func (m *Middleware) Authentication(c *fiber.Ctx) error {
	userSession := getSession(c, m.Session)

	log.Println("User:", userSession, "is logged in")
	log.Println(c.Cookies("session_id"))

	uri := string(c.Request().URI().Path())

	if uri != "/auth/logout" {
		if userSession != nil {
			return c.JSON(utils.ErrorResponse(
				errors.New("Unable to authenticate user"),
				"User is already logged in",
			))
		}
	}

	return c.Next()
}

func getSession(c *fiber.Ctx, store *session.Store) interface{} {
	sess, err := store.Get(c)
	if err != nil {
		return nil
	}

	return sess.Get("username")
}
