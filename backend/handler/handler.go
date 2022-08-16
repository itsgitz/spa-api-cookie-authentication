package handler

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	Session *session.Store
}

func New() *Handler {
	var store = session.New(session.Config{
		// CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
	})

	h := &Handler{}
	h.Session = store

	return h
}
