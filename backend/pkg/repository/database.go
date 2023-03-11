package repository

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func ConnectDatabase() {
	// Values set from .env
	username := "postgres"
	password := "postgres"
	host := "localhost"
	port := 5432
	database := "tolder"
	connection := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		database,
	)
	// Connect to database
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

}

func validateConnection(*sql.DB) {

}
