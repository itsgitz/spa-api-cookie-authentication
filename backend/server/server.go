package server

import (
	"fmt"
	"go-api/handler"
	"go-api/middleware"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var port string

func RunServer() {
	port = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	app := fiber.New(serverConfiguration())
	app.Use(logger.New())
	app.Use(middleware.AuthenticationMiddleware)

	app.Get("/", handler.Home)
	app.Post("/login", handler.Login)

	app.Listen(port)
}

func serverConfiguration() fiber.Config {
	return fiber.Config{
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		EnablePrintRoutes: true,
	}
}
