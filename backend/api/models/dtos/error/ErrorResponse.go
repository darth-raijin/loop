package errorDto

import (
	"time"
)

// Declaring specific DTOs for DomainErrors
var CreateEvent = DomainError{
	DomainErrorCode: 1,
	Message:         "%s is unknown",
}

var EmailNotUnique = DomainError{
	DomainErrorCode: 2,
	Message:         "Email is already in use",
}

var PasswordNotSecure = DomainError{
	DomainErrorCode: 3,
	Message:         "Password not secure",
}

var UsernameInUse = DomainError{
	DomainErrorCode: 4,
	Message:         "Username is already in use",
}

type DomainError struct {
	DomainErrorCode int    `json:"domainErrorCode,omitempty"`
	Message         string `json:"message"`
}

type DomainErrorWrapper struct {
	Timestamp time.Time     `json:"timestamp"`
	Errors    []DomainError `json:"errors"`
}

func (wrapper *DomainErrorWrapper) AddDomainError(error DomainError) *DomainErrorWrapper {
	wrapper.Errors = append(wrapper.Errors, error)
	return wrapper
}
