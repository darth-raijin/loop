package main

import "github.com/gofiber/fiber/v2"

var (
	Version string = "0.0.1"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
