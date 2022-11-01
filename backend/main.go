package main

import (
	_ "github.com/create-go-app/fiber-go-template/docs"
	"github.com/darth-raijin/borealis/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

var (
	Version string = "0.0.1"
)

// @title Fdasdsadsa
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	// Controller declarations
	controllers.Swagger(app) // Register a route for API Docs (Swagger).

	app.Listen(":8080")
}
