package entities

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID uuid.UUID `gorm:"primary_key"`

	Rating int

	// Mapping User as FeedbackGiver
	FeedbackGiverID string // FK for Host
	FeedbackGiver   User   `gorm:"references:Username"`

	// Mapping User
	EventID string // FK for Host
	Event   User

	// Mapping Host User{}
	HostUsername string // FK for Host
	Host         User   `gorm:"references:Username"`

	// Mapping Host User{}
	QuestionID uuid.UUID // FK for Host
	Question   User

	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time
}
