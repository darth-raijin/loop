package controllers

import (
	"net/http"

	"github.com/darth-raijin/borealis/api/models/dtos"
	"github.com/darth-raijin/borealis/api/models/dtos/event"
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
	return c.Status(http.StatusServiceUnavailable).JSON(dtos.ErrorResponse{
		DomainErrorCode: 5001,
		Message:         "No worky worky",
	})
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
	payload := new(event.CreateEventDto)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	validationError := validator.New().Struct(payload)
	if validationError != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(validationError)
	}

	event, err := service.CreateEvent(payload)

	// TODO revisit
	if err == (dtos.ErrorResponse{}) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(event)
}
