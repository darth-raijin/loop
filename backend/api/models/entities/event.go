package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	CreatedBy   string
	Name        string
	Description string
	City        string
	Country     string
	Created     time.Time
	Updated     time.Time

	// Mapping Host User{}
	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
}
