package handlers

import (
	"net/http"

	"github.com/darth-raijin/borealis/api/models/dtos"
	"github.com/gofiber/fiber/v2"
)

// Get a specific event by ID
// @Summary Get a specific event by ID
// @Description Get a specific event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} dtos.Response{data=dtos.Event}
// @Failure 404 {object} dtos.Response{}
// @Router /api/v1/events [get]
func GetEventById(c *fiber.Ctx) error {
	return c.Status(http.StatusServiceUnavailable).JSON(dtos.Response{
		DomainErrorCode: 5001,
		Message:         "No worky worky",
		Data:            nil,
	})
}

// Creates an event
// @Summary Creates an event
// @Description Creates an event
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} dtos.Response{data=dtos.Event}
// @Failure 404 {object} dtos.Response{}
// @Failure 419 {object} dtos.Response{}
// @Router /api/v1/events [get]
func CreateEvent(c *fiber.Ctx) error {
	return c.Status(http.StatusServiceUnavailable).JSON(dtos.Response{
		DomainErrorCode: 5002,
		Message:         "No worky worky",
		Data:            nil,
	})
}
