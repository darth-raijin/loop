package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darth-raijin/borealis/api/routes"
	_ "github.com/darth-raijin/borealis/docs"
	"github.com/joho/godotenv"
)

var (
	Version string = "0.0.1"
)

// @title Borealis
// @version 2.0
// @description REST API server for Borealis aka 'the Feedback' app
func main() {

	app := routes.Initialize()

	fmt.Println(`
	__                               
	/  |                              
	$$ |  ______    ______    ______  
	$$ | /      \  /      \  /      \ 
	$$ |/$$$$$$  |/$$$$$$  |/$$$$$$  |
	$$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |
	$$ |$$ \__$$ |$$ \__$$ |$$ |__$$ |
	$$ |$$    $$/ $$    $$/ $$    $$/ 
	$$/  $$$$$$/   $$$$$$/  $$$$$$$/  
							$$ |      
							$$ |      
							$$/       
	
	`)

	fmt.Println(fmt.Sprintf("BUILD ENVIRONMENT: %s", GetEnvironmentVariable("ENV")))
	app.Listen(":5000")
}

func GetEnvironmentVariable(key string) string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

	return os.Getenv(key)
}
