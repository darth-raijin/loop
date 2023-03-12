package entities

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Question string    `json:"name""`
	Created  time.Time `gorm:"autoCreateTime"`
	Updated  time.Time
}
