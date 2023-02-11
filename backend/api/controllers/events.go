package controllers

import (
	"fmt"
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
// @Param id query string true "id for fetching given event"
// @Failure 404 {object} dtos.Response{}
// @Router /api/v1/events/{id} [get]
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
// @Param Event body dtos.Event true "Required parameters to create an event"
// @Success 200 {object} dtos.Event{}
// @Failure 404 {object} dtos.Response{}
// @Failure 419 {object} dtos.Response{}
// @Router /api/v1/events [post]
func CreateEvent(c *fiber.Ctx) error {
	payload := new(dtos.Event)
	if err := c.BodyParser(payload); err != nil {
		fmt.Println(err)
		return err
	}

	return c.Status(http.StatusServiceUnavailable).JSON(dtos.Response{
		DomainErrorCode: 5002,
		Message:         "No worky worky",
		Data:            nil,
	})
}
