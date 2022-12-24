package dtos

import (
	"github.com/google/uuid"
)

type Event struct {
	id           uuid.UUID
	name         string
	questions    []Question
	participants []string
}
