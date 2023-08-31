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
        "/api/sections": {
            "post": {
                "description": "create new section",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Create new section",
                "operationId": "create-section",
                "parameters": [
                    {
                        "description": "section (slug) name",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.SectionRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Section created",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "message: Failed",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete section",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections"
                ],
                "summary": "Delete section",
                "operationId": "delete-section",
                "parameters": [
                    {
                        "description": "section (slug) name",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.SectionRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Section deleted",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "message: Failed",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/api/sections/users/:id": {
            "get": {
                "description": "get user sections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections-users"
                ],
                "summary": "Get user sections",
                "operationId": "get-user-sections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "description": "add or delete user in sections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sections-users"
                ],
                "summary": "Add user in section",
                "operationId": "add-user",
                "parameters": [
                    {
                        "description": "list of sections to add or delete",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Success",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "message: Failed",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "post": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "user's name",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: User created",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "message: Failed",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Data": {
            "type": "object",
            "properties": {
                "sectionsAdd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "sectionsDelete": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handler.SectionRequestBody": {
            "type": "object",
            "properties": {
                "slug": {
                    "type": "string"
                }
            }
        },
        "handler.UserRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito Proj API",
	Description:      "API Server for AvitoProj Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
