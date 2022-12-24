package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Events(app *fiber.App) {
	event := app.Group("/event")

	// @Param request body main.MyHandler.request true "query params"
	// @Success 200 {object} dtos.Test
	// @Router /test [post]
	event.Get("*", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Placeholder for %s", c.BaseURL()))
	})

	event.Post("*", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Placeholder for %s", c.BaseURL()))
	})
}
