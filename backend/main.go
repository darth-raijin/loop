package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/darth-raijin/borealis/api/routes"
	_ "github.com/darth-raijin/borealis/docs"
	"github.com/joho/godotenv"
)

var (
	Version string = "1.0.0"
)

// @title Loop
// @description REST API server for Loop aka 'the Feedback' app
func main() {

	app := routes.Initialize()

	app.Listen(":8080")
}

func GetEnvironmentVariable(key string) string {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv(key)
}

func printLogo() {
	content, err := ioutil.ReadFile("./resources/logo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
