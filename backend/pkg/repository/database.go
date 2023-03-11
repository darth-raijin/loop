package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func ConnectDatabas2e() {
	// Values set from .env

	// Connect to database
	once.Do(
		func() {
			var err error
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

			db, err = sql.Open(connection)			
		}
	)
	connection, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	db = connection

}

func validateConnection(db *sql.DB) bool {
	if err := db.Ping(); err != nil {
		return false
	}

	return false
}

func ConnectDatabase() {
    once.Do(func() {
		var err error
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

        db, err = sql.Open(connection)
        if err != nil {
            log.Fatal(err)
        }
        if err = db.Ping(); err != nil {
            log.Fatal(err)
        }
        fmt.Println("Connected to the database")
    })
    return db
}