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
        "contact": {
            "name": "Dmitrii Kozii",
            "email": "monrevil@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment/": {
            "get": {
                "description": "get list of comments comments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "List comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Comment"
                            }
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update by json comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update an comment",
                "parameters": [
                    {
                        "description": "Update comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add by json comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Add an comment",
                "parameters": [
                    {
                        "description": "Add comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewComment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/comment/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Show an comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by comment ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete an comment",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "comment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                }
            }
        },
        "/home/": {
            "get": {
                "description": "Use provided link to log in with Google/Facebook",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log in with OAUTH2",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Comment"
                        }
                    }
                }
            }
        },
        "/login/": {
            "get": {
                "description": "Login with app if you have registered with us",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login with the app",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            }
        },
        "/post/": {
            "get": {
                "description": "get list of posts posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "List posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update by json post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Update a post",
                "parameters": [
                    {
                        "description": "Update post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "add by json post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Add a post",
                "parameters": [
                    {
                        "description": "Add post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewPost"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "get": {
                "description": "get post by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Show a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by post ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Delete a post",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.Post"
                        }
                    }
                }
            }
        },
        "/signup/": {
            "post": {
                "description": "Register: email should be unique, pasword from 6 to 25 char long",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signup with app",
                "parameters": [
                    {
                        "description": "Register user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.Comment": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "postID": {
                    "type": "integer"
                }
            }
        },
        "model.NewComment": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "example": "I like tea"
                },
                "name": {
                    "type": "string",
                    "example": "jon"
                },
                "postID": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "model.NewPost": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string",
                    "example": "Tea is good for your health"
                },
                "title": {
                    "type": "string",
                    "example": "Tea"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0",
	Host:        "localhost:1323",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "App Example API",
	Description: "This is a sample api server.",
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
