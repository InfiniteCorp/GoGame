// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/games": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Games"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetAllGamesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/games/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Games"
                ],
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
                            "$ref": "#/definitions/api.GetGameByIDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": ".",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.GetAllUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
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
                            "$ref": "#/definitions/api.ResponseGetUserByID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.GetAllGamesResponse": {
            "type": "object",
            "properties": {
                "games": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/types.GameDefinition"
                    }
                }
            }
        },
        "api.GetAllUsersResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/types.User"
                    }
                }
            }
        },
        "api.GetGameByIDResponse": {
            "type": "object",
            "properties": {
                "game": {
                    "type": "object",
                    "$ref": "#/definitions/types.GameDefinition"
                }
            }
        },
        "api.RegisterRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.RegisterResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/types.User"
                }
            }
        },
        "api.ResponseGetUserByID": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/types.User"
                }
            }
        },
        "types.GameDefinition": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "libPath": {
                    "type": "string"
                }
            }
        },
        "types.User": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
