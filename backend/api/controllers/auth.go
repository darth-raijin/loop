package controllers

import (
	"net/http"

	errorDto "github.com/darth-raijin/loop/api/models/dtos/error"
	"github.com/gofiber/fiber/v2"
)

// Register a User
// @Summary Registers a user
// @Description DomainErrorCodes:
// @Description 2: Email is already in use
// @Description 3: Password not secure
// @Tags Auth
// @Accept json
// @Produce json
// @Failure 422 {object} errorDto.DomainErrorWrapper{}
// @Router /api/auth/register [POST]
func RegisterUser(c *fiber.Ctx) error {

	return c.Status(http.StatusServiceUnavailable).JSON(errorDto.DomainErrorWrapper{})
}
