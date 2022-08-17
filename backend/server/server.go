package server

import (
	"fmt"
	"go-api/handler"
	"go-api/middleware"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	Port        string
	FiberConfig fiber.Config
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	s.setPort()
	s.setConfiguration()

	// Define our App and handlers
	app := fiber.New(s.FiberConfig)
	handlers := handler.New()

	// Pass the session to middlewares
	middlewares := middleware.New(handlers.Session)

	// Use logger middleware
	app.Use(logger.New())

	// Use recover middleware for handle panic
	app.Use(recover.New())

	api := app.Group("/api")
	api.Use(middlewares.VerifyAuthentication)
	api.Get("/", handlers.Home)
	api.Get("/users", handlers.GetUsers)
	api.Get("/notes", handlers.GetNotes)
	api.Post("/notes", handlers.CreateNotes)

	auth := app.Group("/auth")
	auth.Use(middlewares.Authentication)
	auth.Post("/login", handlers.Login)
	auth.Get("/logout", handlers.Logout)

	log.Fatal(app.Listen(s.Port))
}

func (s *Server) setConfiguration() {
	s.FiberConfig.ReadTimeout = 60 * time.Second
	s.FiberConfig.WriteTimeout = 60 * time.Second
	s.FiberConfig.IdleTimeout = 60 * time.Second
	//s.FiberConfig.EnablePrintRoutes = true
}

func (s *Server) setPort() {
	s.Port = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}
