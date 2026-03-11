package main

import (
	"fmt"
	"log"
	"os"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
)

// @title           Vet Shifter API
// @version         1.0
// @description     API for veterinary clinics and shifters.

// @host      localhost:8080
// @BasePath  /

func main() {
	if !utils.LoadEnvironment() {
		log.Println("Warning: .env file not found in common locations, using system environment variables")
	}

	db := database.GetPostgresConnection()
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	router := setupRouter()
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
