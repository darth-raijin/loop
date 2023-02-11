package api

import (
	"github.com/darth-raijin/borealis/api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
)

func Initialize() *fiber.App {
	app := fiber.New()

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		Title: "Borealis",
	}))

	api := app.Group("/api")
	api.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${white} API log \n",
		TimeZone: "Europe/Copenhagen",
	}))

	// Auth endpoints
	auth := api.Group("/auth")
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}]${white} Maraluxa log \n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	auth.Post("/login", controllers.GetEventById)
	auth.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] auth log}\n",
		TimeZone: "Europe/Copenhagen",
	}))

	auth.Get("/register", controllers.GetEventById)
	auth.Post("/register", controllers.GetEventById)

	// Profile
	profile := api.Group("/profile")
	profile.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	profile.Put("/profile", controllers.GetEventById)
	profile.Delete("/profile", controllers.GetEventById)

	return app
}
