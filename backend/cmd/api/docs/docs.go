// Package docs Vet Shifter API Swagger documentation.
// Run `swag init -g cmd/api/main.go --parseDependency` from backend/ to regenerate.
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "schemes": {{ marshal .Schemes }},
  "swagger": "2.0",
  "info": {
    "description": "{{escape .Description}}",
    "title": "{{.Title}}",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "API Support",
      "url": "http://www.vetshifter.io/support"
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
    "/companies": {
      "post": {
        "description": "Creates a company and its owner account. Requires LGPD consent.",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "summary": "Register a new company with owner",
        "operationId": "RegisterCompany",
        "tags": ["companies"],
        "parameters": [
          {
            "description": "Company and owner data",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/controllers.RegisterCompanyRequest" }
          }
        ],
        "responses": {
          "201": {
            "description": "Created with company_id",
            "schema": { "$ref": "#/definitions/controllers.RegisterCompanyResponse" }
          },
          "400": {
            "description": "Invalid request body or validation error",
            "schema": { "$ref": "#/definitions/controllers.ErrorResponse" }
          },
          "409": {
            "description": "CNPJ or email already exists",
            "schema": { "$ref": "#/definitions/controllers.ErrorResponse" }
          },
          "500": {
            "description": "Internal server error",
            "schema": { "$ref": "#/definitions/controllers.ErrorResponse" }
          }
        }
      }
    }
  },
  "definitions": {
    "controllers.RegisterCompanyRequest": {
      "type": "object",
      "required": ["cnpj", "company_name", "owner_name", "email", "phone", "password", "consent_lgpd"],
      "properties": {
        "cnpj": { "type": "string" },
        "company_name": { "type": "string" },
        "owner_name": { "type": "string" },
        "email": { "type": "string", "format": "email" },
        "phone": { "type": "string" },
        "password": { "type": "string" },
        "street": { "type": "string" },
        "number": { "type": "string" },
        "city": { "type": "string" },
        "state": { "type": "string" },
        "zip_code": { "type": "string" },
        "consent_lgpd": { "type": "boolean" }
      }
    },
    "controllers.RegisterCompanyResponse": {
      "type": "object",
      "properties": {
        "company_id": { "type": "string" }
      }
    },
    "controllers.ErrorResponse": {
      "type": "object",
      "properties": {
        "code": { "type": "string" },
        "error": { "type": "string" }
      }
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Vet Shifter API",
	Description:      "API for veterinary clinics and shifters.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
