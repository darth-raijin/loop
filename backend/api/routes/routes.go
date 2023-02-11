package routes

import (
	"github.com/darth-raijin/borealis/api/controllers"
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

	initializeAuth(api)
	// Auth endpoints

	// Profile endpoints
	profile := api.Group("/profile")
	profile.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${red} ${status}${white} - ${method} ${url}  \n",
		TimeZone: "Europe/Copenhagen",
	}))
	profile.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	profile.Put("/profile", controllers.GetEventById)
	profile.Delete("/profile", controllers.GetEventById)

	return app
}
