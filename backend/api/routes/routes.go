package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
)

func Initialize() *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title: "Borealis",
	})).Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))

	api := app.Group("/api")

	// Registering endpoints
	initializeAuth(api)
	initializeEvent(api)
	initializeFeedback(api)
	initializeProfile(api)

	return app
}
