package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Events(app *fiber.App) {
	event := app.Group("/event")

	//	@Summary		Fetches a specific event and the listed questions
	//	@Description	get string by ID
	//	@ID				get-string-by-int
	//	@Tags			bottles
	//	@Accept			json
	//	@Produce		json
	//	@Param			id	path		int	true	"Event ID"
	//	@Success		200	{object}	dtos.Event
	//	@Failure		400	{object}	httputil.HTTPError
	//	@Failure		404	{object}	httputil.HTTPError
	//	@Failure		500	{object}	httputil.HTTPError
	//	@Router			/bottles/{id} [get]
	event.Get("*", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Placeholder for %s", c.BaseURL()))
	})

	event.Post("*", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Placeholder for %s", c.BaseURL()))
	})
}
