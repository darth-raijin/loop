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
	HostUsername string // FK for Host
	Host         User   `gorm:"references:Username"`

	// Mapping Participants
	Participants []*User `gorm:"many2many:event_participants"`

	// Mapping Questions
	Questions []*Question `gorm:"many2many:event_questions"`

	// Mapping Participants
	Feedback []*Feedback `gorm:"many2many:event_feedback"`
}
