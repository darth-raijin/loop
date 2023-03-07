package repository

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConnectDatabase() {
	connection := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

}

func validateConnection(*sql.DB) {

}
