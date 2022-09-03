package server

import (
	"go-api/handler"
	"go-api/middleware"
)

func (s *Server) setRouter() {
	// Call all handlers
	handlers := handler.New()

	// Pass the session from handlers to middlewares
	middlewares := middleware.New(handlers.Session)

	api := s.App.Group("/api")
	api.Use(middlewares.VerifyAuthentication)
	api.Get("/", handlers.Home)
	api.Get("/users", handlers.GetUsers)
	api.Get("/notes", handlers.GetNotes)
	api.Post("/notes", handlers.CreateNotes)

	auth := s.App.Group("/auth")
	auth.Use(middlewares.Authentication)
	auth.Post("/login", handlers.Login)
	auth.Get("/is_login", handlers.IsAuthenticated)
	auth.Get("/logout", handlers.Logout)
}
