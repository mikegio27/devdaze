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
        "/experience": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Resume experience summary",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Experience"
                        }
                    }
                }
            }
        },
        "/objective": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Resume objective",
                "responses": {
                    "200": {
                        "description": "Objective",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/tech-skills": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Resume tech skills summary",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/services.TechSkills"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "services.Experience": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "duties": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "services.TechSkills": {
            "type": "object",
            "properties": {
                "cicd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "databases": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "iac": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "linux": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "monitoring": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "platforms": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.devdaze.org",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "DevDaze API",
	Description:      "API currently serving my resume",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
