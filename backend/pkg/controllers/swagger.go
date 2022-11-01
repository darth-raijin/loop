package controllers

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

func Swagger(app *fiber.App) {
	api := app.Group("/swagger")

	api.Get("*", swagger.HandlerDefault)
}
