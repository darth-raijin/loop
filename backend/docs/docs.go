// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Event"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
                        }
                    },
                    "419": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/dtos.Response"
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
                            "$ref": "#/definitions/dtos.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
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
        "dtos.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "domainErrorCode": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Borealis",
	Description:      "REST API server for Borealis aka 'the Feedback' app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
