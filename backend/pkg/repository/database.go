package repository

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	gormDB *gorm.DB
)

func GormConnectDatabase() {
	once.Do(func() {
		var err error
		username := "postgres"
		password := "postgres"
		host := "localhost"
		port := 5432
		database := "loop"
		connection := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
			username,
			password,
			host,
			port,
			database,
		)

		gormDB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	})

	migrateEntities()
}

func migrateEntities() {

}
