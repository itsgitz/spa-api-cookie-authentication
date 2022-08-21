package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	Port        string
	App         *fiber.App
	FiberConfig fiber.Config
	CORS        cors.Config
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() {
	s.setPort()
	s.setConfiguration()

	// Define our App and handlers
	s.App = fiber.New(s.FiberConfig)

	// Use logger middleware
	s.App.Use(logger.New())

	// Use recover middleware for handle panic
	s.App.Use(recover.New())

	// CORS setting
	s.setCors()
	s.App.Use(cors.New(s.CORS))

	// Set router
	s.setRouter()

	// Listen ...
	log.Fatal(s.App.Listen(s.Port))
}

func (s *Server) setConfiguration() {
	s.FiberConfig.ReadTimeout = 60 * time.Second
	s.FiberConfig.WriteTimeout = 60 * time.Second
	s.FiberConfig.IdleTimeout = 60 * time.Second
}

func (s *Server) setPort() {
	s.Port = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}
