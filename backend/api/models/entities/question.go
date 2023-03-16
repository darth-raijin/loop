package entities

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID       uuid.UUID
	Question string
	Created  time.Time
	Updated  time.Time
}
