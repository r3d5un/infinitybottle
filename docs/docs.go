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
        "/v1/contribution": {
            "get": {
                "description": "List all infinity bottles contributions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contribution"
                ],
                "summary": "List all infinity bottle contributions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "brand name to search for",
                        "name": "brandName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "tags to search for",
                        "name": "tags",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.Contribution"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/contributions": {
            "post": {
                "description": "Add a new contribution to an infinity bottle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contribution"
                ],
                "summary": "Add a new contribution to an infinity bottle",
                "parameters": [
                    {
                        "description": "New contribution to an infinity bottle",
                        "name": "Contribution",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ContributionPost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.ContributionPost"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            }
        },
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
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Update all information about an infinity bottle contribution by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contribution"
                ],
                "summary": "Update an infinity bottle contribution by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update contribution to an infinity bottle",
                        "name": "Contribution",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ContributionPost"
                        }
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
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an infinity bottle contribution by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contribution"
                ],
                "summary": "Delete an infinity bottle contribution by ID",
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
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
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
                            "$ref": "#/definitions/main.HealthCheckMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/v1/infinitybottles": {
            "get": {
                "description": "List all infinity bottles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infinityBottle"
                ],
                "summary": "List all infinity bottles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bottle name to search for",
                        "name": "bottleName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/data.InfinityBottle"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new infinity bottle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infinityBottle"
                ],
                "summary": "Create a new infinity bottle",
                "parameters": [
                    {
                        "description": "New infinity bottle",
                        "name": "InfinityBottle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.InfinityBottlePost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.InfinityBottlePost"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
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
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Update all information about an infinity bottle by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infinityBottle"
                ],
                "summary": "Update an infinity bottle by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update to an infinity bottle",
                        "name": "InfinityBottle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.InfinityBottlePost"
                        }
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
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an infinity bottle by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infinityBottle"
                ],
                "summary": "Delete an infinity bottle by ID",
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
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorMessage"
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
                },
                "version": {
                    "type": "integer"
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
                },
                "version": {
                    "type": "integer"
                }
            }
        },
        "main.ContributionPost": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "brandName": {
                    "type": "string"
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
        "main.ErrorMessage": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "main.HealthCheckMessage": {
            "type": "object",
            "properties": {
                "environment": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "main.InfinityBottlePost": {
            "type": "object",
            "properties": {
                "bottleName": {
                    "type": "string"
                },
                "emptyStart": {
                    "type": "boolean"
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
