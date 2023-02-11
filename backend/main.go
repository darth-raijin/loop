package main

import (
	"github.com/darth-raijin/borealis/api"
	_ "github.com/darth-raijin/borealis/docs"
)

var (
	Version string = "0.0.1"
)

// @title Borealis
// @version 2.0
// @description REST API server for Borealis aka 'the Feedback' app
func main() {
	app := api.Initialize()

	// Registering controllers
	app.Listen(":5000")
}
