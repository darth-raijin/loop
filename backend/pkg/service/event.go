package service

import (
	"time"

	"github.com/darth-raijin/borealis/api/models/dtos"
	"github.com/darth-raijin/borealis/api/models/dtos/event"
	"github.com/google/uuid"
)

func CreateEvent(payload *event.CreateEventDto) (event.FetchEventDto, dtos.ErrorResponse) {
	sampleUUID, _ := uuid.NewUUID()

	// DB transaction here -> pass err from here to return err

	if 2 == 4 {
		return event.FetchEventDto{}, dtos.ErrorResponse{
			Message:   "Error creating event",
			Timestamp: time.Now().UTC(),
		}
	}

	return event.FetchEventDto{
		ID:          sampleUUID,
		Name:        payload.Name,
		Description: payload.Description,
		City:        payload.City,
		Country:     payload.Country,
	}, dtos.ErrorResponse{}
}
