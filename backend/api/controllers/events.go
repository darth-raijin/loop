package controllers

import (
	"net/http"
	"time"

	"github.com/darth-raijin/borealis/api/models/dtos"
	createEventDto "github.com/darth-raijin/borealis/api/models/dtos/event/createevent"
	"github.com/darth-raijin/borealis/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Get a specific event by ID
// @Summary Get a specific event by ID
// @Description Get a specific event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id query string true "id for fetching given event"
// @Failure 404 {object} dtos.ErrorResponse{}
// @Router /api/v1/events/{id} [get]
func GetEventById(c *fiber.Ctx) error {
	return c.Status(http.StatusServiceUnavailable).JSON(dtos.DomainErrorWrapper{})
}

// Creates an event
// @Summary Creates an event
// @Description Creates an event
// @Tags Event
// @Accept json
// @Produce json
// @Param Event body dtos.Event true "Required parameters to create an event"
// @Success 201 {object} dtos.CreateEventDto{}
// @Failure 419 {object} dtos.ErrorResponse{}
// @Router /api/v1/events [post]
func CreateEvent(c *fiber.Ctx) error {
	payload := new(createEventDto.CreateEventRequest)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	validationError := validator.New().Struct(payload)
	if validationError != nil {
		// Check list and create a DomainErrorWrapper and return
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dtos.DomainError{
			Message:   validationError.Error(),
			Timestamp: time.Now().UTC(),
		})
	}

	event, err := service.CreateEvent(payload)

	// TODO revisit
	if _ != (dtos.DomainError{}) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(event)
}
