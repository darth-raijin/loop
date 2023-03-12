package service

import (
	errorDto "github.com/darth-raijin/loop/api/models/dtos/error"
	createEventDto "github.com/darth-raijin/loop/api/models/dtos/event/createevent"
	"github.com/darth-raijin/loop/api/models/entities"
	"github.com/darth-raijin/loop/pkg/repository"
)

func CreateEvent(payload *createEventDto.CreateEventRequest) (createEventDto.CreateEventResponse, errorDto.DomainErrorWrapper) {
	entity := entities.Event{
		Name:        "Sample Event",
		Description: "Sample Description",
		City:        "Copenhagen",
		Country:     "Denmark",
		Questions:   []entities.Question{},
	}
	// DB transaction here -> pass err from here to return err
	repository.GormDB.Create(&entity)

	return createEventDto.CreateEventResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		City:        entity.City,
		Country:     entity.Country,
	}, errorDto.DomainErrorWrapper{}
}
