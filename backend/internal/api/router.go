package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	authcontrollers "rodrigoorlandini/vet-shifter/internal/auth/infrastructure/controllers"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/controllers"
	shiftcontrollers "rodrigoorlandini/vet-shifter/internal/shifts/infrastructure/controllers"
	shiftvetcontrollers "rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	router.POST("/auth/login", authcontrollers.NewLoginController().Handle)
	router.POST("/companies", controllers.NewRegisterCompanyController().Handle)
	router.POST("/shift-vets", shiftvetcontrollers.NewRegisterShiftVetController().Handle)
	router.POST("/shifts", shiftcontrollers.NewShiftController().Create)
	router.GET("/shifts", shiftcontrollers.NewShiftController().List)
	return router
}
