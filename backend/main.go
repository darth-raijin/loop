package main

import (
	_ "github.com/darth-raijin/borealis/docs"
	"github.com/gofiber/fiber/v2"
)

var (
	Version string = "0.0.1"
)

// @title Borealis
// @version 2.0
// @description REST API server for Borealis aka 'the Feedback' app
func main() {
	app := fiber.New()

	// Registering controllers
	controllers.Swagger(app)

	app.Listen(":8080")
}
