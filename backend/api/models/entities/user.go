package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string `gorm:"unique;not null"`
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string // Do some crazy HASH + SALT
	Created  time.Time
}
