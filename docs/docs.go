// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/send": {
            "post": {
                "description": "Send plain text email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Email"
                ],
                "summary": "Send plain text email",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.PlainTextEmail"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/send/{slug}": {
            "post": {
                "description": "Send template email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Email"
                ],
                "summary": "Send template email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Template slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.TemplateEmail"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "types.PlainTextEmail": {
            "type": "object",
            "required": [
                "subject",
                "to"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "types.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "types.TemplateEmail": {
            "type": "object",
            "required": [
                "subject",
                "to"
            ],
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "subject": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8082",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Demo SMTP API",
	Description:      "This is a demo SMTP API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
