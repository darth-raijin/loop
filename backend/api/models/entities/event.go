package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	City        string     `json:"city,omitempty"`
	Country     string     `json:"country,omitempty"`
	Questions   []Question `json:"questions,omitempty" validate:"required"`
	Created     time.Time  `gorm:"autoCreateTime"`
	Updated     time.Time
}
