package event

import (
	"github.com/google/uuid"
)

type FetchEventDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	City        string    `json:"city,omitempty"`
	Country     string    `json:"country,omitempty"`
}
