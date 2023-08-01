// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/contributions/{id}": {
            "get": {
                "description": "Retrieve all information about an infinity bottle contribution by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contribution"
                ],
                "summary": "Get an infinity bottle contribution by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.Contribution"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/healthcheck": {
            "get": {
                "description": "Perform a basic request to check if the service is available",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Basic healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/infinitybottles/{id}": {
            "get": {
                "description": "Retrieve all information about an infinity bottle by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infinityBottle"
                ],
                "summary": "Get an infinity bottle by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/data.InfinityBottle"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "data.Contribution": {
            "type": "object",
            "properties": {
                "addedAt": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                },
                "brandName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "infinityBottleID": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "data.InfinityBottle": {
            "type": "object",
            "properties": {
                "bottleName": {
                    "type": "string"
                },
                "contributions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/data.Contribution"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "emptyStart": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "numberOfContributions": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Infinity Bottle API",
	Description:      "This is a REST API built to keep track of whisky infinity bottles and their history",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
