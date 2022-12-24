package routes

import (
	"github.com/darth-raijin/borealis/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
)

func Initialize() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeZone: "Europe/Copenhagen",
	}))

	app.Get("*", swagger.New(swagger.Config{ // custom
		Title: "Borealis",
	}))

	api := app.Group("/api")

	// Initializing v1 endpoints
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "Borealis v1",
		})
		return c.Next()
	})

	// Event endpoints
	v1.Get("/events/:id", handlers.GetEventById)
	v1.Post("/events/", handlers.CreateEvent)

	return app
}
