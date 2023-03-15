// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/register": {
            "post": {
                "description": "DomainErrorCodes:\n2: Email is already in use\n3: Password not secure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Registers a user",
                "responses": {
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errorDto.DomainErrorWrapper"
                        }
                    }
                }
            }
        },
        "/api/v1/events": {
            "post": {
                "description": "Creates an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Creates an event",
                "parameters": [
                    {
                        "description": "Required parameters to create an event",
                        "name": "Event",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Event"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/createEventDto.CreateEventRequest"
                        }
                    },
                    "419": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errorDto.DomainErrorWrapper"
                        }
                    }
                }
            }
        },
        "/api/v1/events/{id}": {
            "get": {
                "description": "Get a specific event by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get a specific event by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id for fetching given event",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorDto.DomainErrorWrapper"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "createEventDto.CreateEventRequest": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "maxLength": 255
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "Sample description"
                },
                "name": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 3
                }
            }
        },
        "dtos.Event": {
            "type": "object",
            "required": [
                "id",
                "name",
                "questions"
            ],
            "properties": {
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "name": {
                    "type": "string",
                    "example": "Go Crash Course"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.Question"
                    }
                }
            }
        },
        "dtos.Question": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "format": "uuid"
                },
                "question": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "errorDto.DomainError": {
            "type": "object",
            "properties": {
                "domainErrorCode": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "errorDto.DomainErrorWrapper": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errorDto.DomainError"
                    }
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Loop",
	Description:      "REST API server for Loop aka 'the Feedback' app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
