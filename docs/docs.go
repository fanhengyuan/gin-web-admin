// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-01 16:17:29.19162062 +0800 CST m=+0.042160491

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "An example of gin",
        "title": "Gin Web Test",
        "termsOfService": "https://github.com/fanhengyuan/gin-test",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/fanhengyuan/gin-test/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "user login",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/userService.AuthStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/report": {
            "post": {
                "description": "User Report Information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Report"
                ],
                "summary": "Report Information",
                "parameters": [
                    {
                        "description": "上报信息",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/reportService.ReportStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/sys/router": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get router list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SYS"
                ],
                "summary": "后端存在路由列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/test/font": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Base64 Decode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "base64 string",
                        "name": "base64",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/test/ping": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Test Ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "p",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page limit",
                        "name": "n",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/change_password": {
            "put": {
                "description": "密码修改",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "description": "user change password",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/userService.ChangePasswordStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/logged_in": {
            "get": {
                "description": "当前登录用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "当前登录用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "put": {
                "description": "用户登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "reportService.ReportStruct": {
            "type": "object",
            "required": [
                "name",
                "phone"
            ],
            "properties": {
                "activity_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "minLength": 1
                },
                "phone": {
                    "type": "string",
                    "minLength": 4
                }
            }
        },
        "userService.AuthStruct": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 4
                },
                "user_name": {
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "userService.ChangePasswordStruct": {
            "type": "object",
            "required": [
                "new_password",
                "old_password"
            ],
            "properties": {
                "new_password": {
                    "type": "string",
                    "minLength": 6
                },
                "old_password": {
                    "type": "string",
                    "minLength": 4
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
