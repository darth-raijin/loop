package main

import (
	"fmt"
	"os"

	"github.com/darth-raijin/borealis/api/routes"
	_ "github.com/darth-raijin/borealis/docs"
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

	fmt.Println(fmt.Sprintf("BUILD ENVIRONMENT: %s", os.Getenv("env")))
	app.Listen(":5000")
}
