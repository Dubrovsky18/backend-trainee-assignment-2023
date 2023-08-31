// Code generated by swaggo/swag. DO NOT EDIT.

package api

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
        "/api/v1/slug/create": {
            "post": {
                "description": "Create a new slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new slug",
                "parameters": [
                    {
                        "description": "Slug object",
                        "name": "slug",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.Slug"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.StatusResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/slug/delete/{name_slug}": {
            "delete": {
                "description": "Delete a slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a slug",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the slug",
                        "name": "name_slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "name_slug",
                        "name": "slug",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.Slug"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.StatusResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/slug/get_all": {
            "get": {
                "description": "Get all slugs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all slugs",
                "responses": {
                    "200": {
                        "description": "List of slugs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.Slug"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/status": {
            "get": {
                "description": "get status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Application Status",
                "operationId": "get-status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_web_controllers_apiv1_status.ResponseDoc"
                        }
                    }
                }
            }
        },
        "/api/v1/users/add_del_slug/{user_id}": {
            "post": {
                "description": "Add or delete slugs for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add or delete slugs for a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "List of slugs to add or delete",
                        "name": "listAddDel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.AddRemoveUserSlug"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.StatusResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/create": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User object",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User id",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/users/delete/{user_id}": {
            "delete": {
                "description": "Delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg.StatusResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/extra/history/{user_id}": {
            "get": {
                "description": "Get segments history for a user within a specified period",
                "produces": [
                    "application/json"
                ],
                "summary": "Get segments history for a user",
                "operationId": "get-segments-history",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User history information",
                        "name": "history",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.UserHistory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "CSV file containing user's segment history",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/get_slugs/{user_id}": {
            "get": {
                "description": "Get user's slugs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user's slugs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User slugs",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_app_build.Info": {
            "type": "object",
            "properties": {
                "arch": {
                    "type": "string"
                },
                "build_date": {
                    "type": "string"
                },
                "commit_hash": {
                    "type": "string"
                },
                "compiler": {
                    "type": "string"
                },
                "go_version": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.AddRemoveUserSlug": {
            "type": "object",
            "properties": {
                "add_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "del_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.Slug": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name_slug": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_models.UserHistory": {
            "type": "object",
            "properties": {
                "month_finish": {
                    "type": "integer"
                },
                "month_start": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                },
                "year_finish": {
                    "type": "integer"
                },
                "year_start": {
                    "type": "integer"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "internal_web_controllers_apiv1_status.Response": {
            "type": "object",
            "properties": {
                "build": {
                    "$ref": "#/definitions/github_com_Dubrovsky18_backend-trainee-assignment-2023_internal_app_build.Info"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "internal_web_controllers_apiv1_status.ResponseDoc": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "attributes": {
                            "$ref": "#/definitions/internal_web_controllers_apiv1_status.Response"
                        }
                    }
                }
            }
        },
        "pkg.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "pkg.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
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
