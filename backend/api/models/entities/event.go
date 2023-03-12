package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string
	Description string
	City        string
	Country     string
	Questions   []Question
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time
}
