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
        "/loan-application": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loan-origination"
                ],
                "summary": "Create New Loan APplication",
                "operationId": "create-loan-application",
                "parameters": [
                    {
                        "description": "Loan Application ",
                        "name": "loanApplicationRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoanApplicationInputStep"
                        }
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
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            }
        },
        "/loan-application/{workflow_id}/{run_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loan-origination"
                ],
                "summary": "GetLoan Application state",
                "operationId": "get-loan-application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workflow ID",
                        "name": "workflow_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Run ID",
                        "name": "run_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.QueryResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.HTTPError": {
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
        "common.QueryResult": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/common.State"
                }
            }
        },
        "common.State": {
            "type": "string",
            "enum": [
                "initialized",
                "submitted",
                "approved",
                "rejected",
                "closed"
            ],
            "x-enum-varnames": [
                "Initialized",
                "Submitted",
                "Approved",
                "Rejected",
                "Closed"
            ]
        },
        "dto.LoanApplicationInputStep": {
            "type": "object",
            "properties": {
                "aadhaar_number": {
                    "type": "string"
                },
                "applicant_name": {
                    "description": "ApplicationNo string ` + "`" + `json:\"application_no\"` + "`" + `",
                    "type": "string"
                },
                "pan_number": {
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
