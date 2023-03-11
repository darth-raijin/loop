package service

import (
	"github.com/darth-raijin/borealis/api/models/dtos"
	createEventDto "github.com/darth-raijin/borealis/api/models/dtos/event/createevent"

	"github.com/google/uuid"
)

func CreateEvent(payload *createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, dtos.DomainErrorWrapper) {
	sampleUUID, _ := uuid.NewUUID()

	// DB transaction here -> pass err from here to return err

	if 2 == 4 {
		return createEventDto.CreateEventResponse{}, dtos.DomainErrorWrapper{}
	}

	return createEventDto.CreateEventResponse{
		ID:          sampleUUID,
		Name:        payload.Name,
		Description: payload.Description,
		City:        payload.City,
		Country:     payload.Country,
	}, dtos.DomainErrorWrapper{}
}
