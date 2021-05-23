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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "aumujun@gmail.com",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/article/get": {
            "get": {
                "description": "文章获取接口",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "文章管理相关接口"
                ],
                "summary": "文章获取接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求页数",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "请求页码",
                        "name": "page_size",
                        "in": "query"
                    }
                ]
            }
        },
        "/article/post": {
            "post": {
                "description": "按照一定规则新建文章",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章管理相关接口"
                ],
                "summary": "新增文章接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章hash值，为空则后端计算",
                        "name": "content_hash",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "description": "是否显示背景图，默认false，为true时必须选择设置random和pic其中之一",
                        "name": "background_visible",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "description": "当background_visible为true时",
                        "name": "background_random",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "当background_visible为true时",
                        "name": "background_pic",
                        "in": "formData"
                    }
                ]
            }
        },
        "/article/remove": {
            "post": {
                "description": "文章删除接口 接收一个文章对象",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "文章管理相关接口"
                ],
                "summary": "文章删除接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "文章对象",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                ]
            }
        },
        "/upload/pic": {
            "post": {
                "description": "图片上传接口 单纯只接收图片",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "上传管理相关接口"
                ],
                "summary": "图片上传接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图片hash 为空则后端进行计算",
                        "name": "file_hash",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "图片",
                        "name": "local_file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.MsgResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录，返回JWT",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理相关接口"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
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
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理相关接口"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "username",
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
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Article": {
            "type": "object",
            "properties": {
                "background_pic": {
                    "type": "string",
                    "example": "http://example.com/a.png"
                },
                "background_random": {
                    "type": "boolean",
                    "example": true
                },
                "background_visible": {
                    "type": "boolean",
                    "example": true
                },
                "category": {
                    "type": "string",
                    "example": "default"
                },
                "content": {
                    "type": "string",
                    "example": "你好，世界！"
                },
                "content_hash": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "example": "世界，你好！"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.MsgResponse": {
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
	Version:     "0.0.1",
	Host:        "api.amujun.com",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "OPiece",
	Description: "博客后台程序REST API接口",
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
