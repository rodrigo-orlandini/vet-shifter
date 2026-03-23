package router

import (
	"rodrigoorlandini/vet-shifter/cmd/api/docs"
	sharedmiddleware "rodrigoorlandini/vet-shifter/internal/_shared/api/middleware"
	authcontrollers "rodrigoorlandini/vet-shifter/internal/auth/infrastructure/controllers"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/middleware"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/controllers"
	veterinarycontrollers "rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.BasePath = "/"
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(sharedmiddleware.CORSMiddleware())
	r.Use(sharedmiddleware.ErrorLoggingMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	companies := r.Group("/companies")
	{
		registerCompany := controllers.NewRegisterCompanyController()
		companies.POST("", registerCompany.Handle)
	}

	veterinaries := r.Group("/veterinaries")
	{
		registerShiftVeterinary := veterinarycontrollers.NewRegisterShiftVeterinaryController()
		veterinaries.POST("", registerShiftVeterinary.Handle)
	}

	auth := r.Group("/auth")
	{
		getUserTypeController := authcontrollers.NewGetUserTypeController()
		auth.GET("/user-type", getUserTypeController.Handle)

		loginCompanyOwnerController := authcontrollers.NewLoginCompanyOwnerController()
		loginVeterinaryController := authcontrollers.NewLoginVeterinaryController()
		auth.POST("/login/owner", loginCompanyOwnerController.Handle)
		auth.POST("/login/veterinary", loginVeterinaryController.Handle)

		forgotPasswordController := authcontrollers.NewForgotPasswordController()
		auth.POST("/forgot-password", forgotPasswordController.Handle)

		resetPasswordController := authcontrollers.NewResetPasswordController()
		auth.POST("/reset-password", resetPasswordController.Handle)

		logoutController := authcontrollers.NewLogoutController()
		auth.POST("/logout", middleware.AuthMiddleware(), logoutController.Handle)
	}

	return r
}
