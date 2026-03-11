package main

import (
	"rodrigoorlandini/vet-shifter/cmd/api/docs"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	companies := r.Group("/companies")
	{
		registerCompany := controllers.NewRegisterCompanyController()
		companies.POST("", registerCompany.Handle)
	}

	return r
}

func init() {
	docs.SwaggerInfo.BasePath = "/"
}
