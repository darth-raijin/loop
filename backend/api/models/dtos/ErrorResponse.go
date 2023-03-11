package dtos

import (
	"time"
)

type DomainError struct {
	DomainErrorCode int       `json:"domainErrorCode,omitempty"`
	Message         string    `json:"message"`
	Timestamp       time.Time `json:"timestamp"`
}

type DomainErrorWrapper struct {
	Errors []DomainError `json:errors"`
}

func (wrapper *DomainErrorWrapper) AddDomainError(error DomainError) *DomainErrorWrapper {
	_ = append(wrapper.Errors, error)
	return wrapper
}
