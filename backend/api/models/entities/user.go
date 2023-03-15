package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	Username      string
	Name          string
	Email         string    // Maybe HASH?
	Password      string    // Do some crazy HASH + SALT
	Hosted        []Event   `gorm:"foreignKey:ID"`
	Participating []Event   `gorm:"foreignKey:ID"`
	Created       time.Time `gorm:"autoCreateTime"`
}
