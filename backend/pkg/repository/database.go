package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/darth-raijin/loop/api/models/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
)

func GormConnectDatabase() {
	var err error

	// Values declared in .env -> keep secret pls
	username := "postgres"
	password := "postgres"
	host := "127.0.0.1"
	port := 5432
	database := "tolder"
	retryDelay := 10

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
		log.Default().Println(fmt.Printf("Failed connecting to database, retrying in %v seconds", retryDelay))
		time.Sleep(time.Second * time.Duration(retryDelay))
		GormConnectDatabase()
	}

	migrateEntities()
}

func migrateEntities() {
	GormDB.AutoMigrate(&entities.Question{})
	GormDB.AutoMigrate(&entities.User{})
	GormDB.AutoMigrate(&entities.Event{})
	GormDB.AutoMigrate(&entities.Feedback{})
}
