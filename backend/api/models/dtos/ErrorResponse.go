package dtos

import (
	"time"
)

type ErrorResponse struct {
	DomainErrorCode int       `json:"domainErrorCode,omitempty"`
	Message         string    `json:"message"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
