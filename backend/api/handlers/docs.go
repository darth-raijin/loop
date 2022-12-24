package handlers

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

// @tags docs
func Swagger(app *fiber.App) {
	api := app.Group("/swagger")

	api.Get("*", swagger.New(swagger.Config{ // custom
		Title: "Borealis",
	}))
}
