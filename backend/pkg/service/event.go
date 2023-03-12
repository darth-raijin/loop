package service

import (
	errorDto "github.com/darth-raijin/borealis/api/models/dtos/error"
	createEventDto "github.com/darth-raijin/borealis/api/models/dtos/event/createevent"

	"github.com/google/uuid"
)

func CreateEvent(payload *createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {
	sampleUUID, _ := uuid.NewUUID()

	// DB transaction here -> pass err from here to return err

	created := createEventDto.CreateEventResponse{
		ID:          sampleUUID,
		Name:        payload.Name,
		Description: payload.Description,
		City:        payload.City,
		Country:     payload.Country,
	}

	return created, errorDto.DomainErrorWrapper{}
}
