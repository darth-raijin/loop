package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/darth-raijin/loop/api/models/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	GormDB *gorm.DB
)

func GormConnectDatabase() {

	var err error
	username := "admin"
	password := "test"
	host := "127.0.0.1"
	port := 5432
	database := "loop"
	connection := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		database,
	)

	GormDB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		log.Default().Println("Failed connecting to database, retrying in 10 seconds")
		GormConnectDatabase()
	}

	migrateEntities()
}

func migrateEntities() {
	GormDB.AutoMigrate(&entities.Event{})
}
