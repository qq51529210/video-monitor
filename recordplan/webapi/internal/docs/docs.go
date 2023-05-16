// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/week_plans": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "列表",
                "parameters": [
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "创建时间，时间戳，范围匹配，CreatedAt \u003e= CreatedAtAfter",
                        "name": "afterCreatedAt",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "更新时间，时间戳，范围匹配，UpdateAt \u003e= UpdatedAtAfter",
                        "name": "afterUpdatedAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "创建时间，时间戳，范围匹配，CreatedAt \u003c CreatedBefore",
                        "name": "beforeCreatedAt",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "更新时间，时间戳，范围匹配，UpdateAt \u003c UpdatedAtBefore",
                        "name": "beforeUpdatedAt",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "条数，小于 1 不匹配",
                        "name": "count",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "description": "是否禁用，精确匹配",
                        "name": "enable",
                        "in": "query"
                    },
                    {
                        "maxLength": 32,
                        "type": "string",
                        "description": "名称，模糊匹配",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "偏移，小于 0 不匹配",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "排序，\"column [desc]\"",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "保存的天数",
                        "name": "saveDays",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.ListData-db_WeekPlan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "添加",
                "parameters": [
                    {
                        "description": "数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/weekplans.postReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/internal.IDResult-int64"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "批量删除",
                "parameters": [
                    {
                        "description": "条件",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/internal.RowResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            }
        },
        "/week_plans/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据库 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.WeekPlan"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "删除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据库 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "周计划"
                ],
                "summary": "修改",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "WeekPlan.ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/weekplans.patchReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/internal.RowResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.ListData-db_WeekPlan": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.WeekPlan"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "db.TimePeroid": {
            "type": "object",
            "properties": {
                "end": {
                    "description": "结束时间戳",
                    "type": "string"
                },
                "start": {
                    "description": "开始时间戳",
                    "type": "string"
                }
            }
        },
        "db.WeekPlan": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "数据库的创建时间，时间戳，",
                    "type": "integer"
                },
                "enable": {
                    "description": "是否禁用\n0: 禁用\n1: 启用",
                    "type": "integer"
                },
                "id": {
                    "description": "数据库ID",
                    "type": "integer"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "peroids": {
                    "description": "是一个 RecordTime 的 JSON 数组",
                    "type": "string"
                },
                "saveDays": {
                    "description": "保存的天数",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "数据库的更新时间",
                    "type": "integer"
                }
            }
        },
        "internal.Error": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "详细信息"
                },
                "phrase": {
                    "description": "简短信息",
                    "type": "string"
                }
            }
        },
        "internal.IDResult-int64": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "数据库 ID",
                    "type": "integer"
                }
            }
        },
        "internal.RowResult": {
            "type": "object",
            "properties": {
                "row": {
                    "description": "行数",
                    "type": "integer"
                }
            }
        },
        "weekplans.patchReq": {
            "type": "object",
            "properties": {
                "enable": {
                    "description": "是否禁用\n0: 禁用\n1: 启用",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ]
                },
                "name": {
                    "description": "名称",
                    "type": "string",
                    "maxLength": 32
                },
                "peroids": {
                    "description": "是一个 RecordTime 的 JSON 数组",
                    "type": "array",
                    "maxItems": 7,
                    "minItems": 1,
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/db.TimePeroid"
                        }
                    }
                },
                "saveDays": {
                    "description": "保存的天数",
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "weekplans.postReq": {
            "type": "object",
            "required": [
                "name",
                "peroids"
            ],
            "properties": {
                "enable": {
                    "description": "是否禁用，默认是 1\n0: 禁用\n1: 启用",
                    "type": "integer",
                    "enum": [
                        0,
                        1
                    ]
                },
                "name": {
                    "description": "名称",
                    "type": "string",
                    "maxLength": 32
                },
                "peroids": {
                    "description": "是一个 RecordTime 的 JSON 数组",
                    "type": "array",
                    "maxItems": 7,
                    "minItems": 1,
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/db.TimePeroid"
                        }
                    }
                },
                "saveDays": {
                    "description": "保存的天数，默认是 1",
                    "type": "integer",
                    "minimum": 1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "接口文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}