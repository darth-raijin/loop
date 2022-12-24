package dtos

type Response struct {
	DomainErrorCode int         `json:"domainErrorCode"`
	Data            interface{} `json:"data"`
	Message         string      `json:"message"`
}
