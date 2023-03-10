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
        "/captcha/image/get": {
            "get": {
                "description": "获取图片验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证码相关"
                ],
                "summary": "获取图片验证码",
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
        "/captcha/image/get/{filename}": {
            "get": {
                "description": "获取历史图片验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证码相关"
                ],
                "summary": "获取历史图片验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图片文件名",
                        "name": "filename",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "/captcha/image/proof": {
            "post": {
                "description": "校验验证码是否正确",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "验证码相关"
                ],
                "summary": "图片验证码校验",
                "parameters": [
                    {
                        "type": "string",
                        "description": "验证码的ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要校验的值",
                        "name": "val",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":\"认证通过\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"status\":\"认证失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/channel/get": {
            "post": {
                "description": "获取频道",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "获取频道",
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
        "/channel/group/get": {
            "post": {
                "description": "获取频道组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "获取频道组",
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
        "/message/get": {
            "post": {
                "description": "获取消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "获取消息",
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
        "/message/send": {
            "post": {
                "description": "发送消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "发送消息",
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
        "/sys/metrics": {
            "get": {
                "description": "获取Prometheus统计的指标",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统相关"
                ],
                "summary": "获取指标",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sys/swagger/index.html": {
            "get": {
                "description": "获取API文档",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统相关"
                ],
                "summary": "获取API文档",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/topic/join": {
            "post": {
                "description": "加入话题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "加入话题",
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
        "/topic/new": {
            "post": {
                "description": "新增话题",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "即时通讯相关"
                ],
                "summary": "新增话题",
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
        "/user/login": {
            "post": {
                "description": "通过RPC对话验证码服务。判断验证码是否复合要求在决定是否放行，网关不鉴别JWT留给用户服务实现，主要怕别人伪造接口，登录成功会返回ws服务器的地址",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱/手机号/用户名",
                        "name": "key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "登录密码",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":\"登录成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"status\":\"登录失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "pwd",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "key",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":\"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"status\":\"注册失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "这是Truffle微服务的网关",
	Description:      "Truffle由好几个部分组成，在网关部分可以访问普罗米修斯",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
