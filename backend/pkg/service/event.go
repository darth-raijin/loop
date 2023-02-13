package service

import (
	"github.com/darth-raijin/borealis/api/models/dtos/event"
	"github.com/google/uuid"
)

func CreateEvent(payload *event.CreateEventDto) (event.FetchEventDto, error) {
	sampleUUID, _ := uuid.NewUUID()

	return event.FetchEventDto{
		ID:          sampleUUID,
		Name:        payload.Name,
		Description: payload.Description,
		City:        payload.City,
		Country:     payload.Country,
	}, nil
}
