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
        "/exoplanets": {
            "get": {
                "description": "This end point returns list of all exoplanetts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exoplanets"
                ],
                "summary": "Get All exoplanete",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exoplanet"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Product information",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Exoplanet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exoplanet"
                        }
                    }
                }
            }
        },
        "/exoplanets/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exoplanets"
                ],
                "summary": "Get Explonets by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Exoplanet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exoplanet"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exoplanets"
                ],
                "summary": "Update Explonets by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Exoplanet ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Exoplanet"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Delete Explonets by ID",
                "responses": {}
            }
        },
        "/ping": {
            "get": {
                "description": "This end point respond to pings",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Exoplanet": {
            "type": "object",
            "required": [
                "description",
                "distance",
                "name",
                "radius"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "distance": {
                    "description": "distance from Earth in light years",
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "mass": {
                    "description": "mass in Earth-mass units, only for Terrestrial",
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "radius": {
                    "description": "radius in Earth-radius units",
                    "type": "number"
                },
                "type": {
                    "description": "GasGiant or Terrestrial",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.ExoplanetType"
                        }
                    ]
                }
            }
        },
        "models.ExoplanetType": {
            "type": "string",
            "enum": [
                "GasGiant",
                "Terrestrial"
            ],
            "x-enum-varnames": [
                "GasGiant",
                "Terrestrial"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
