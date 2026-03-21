package main

import (
	"fmt"
	"log/slog"

	"rodrigoorlandini/vet-shifter/internal/_shared/api/router"
	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/logger"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
)

// @title           Vet Shifter API
// @version         1.0
// @description     API for veterinary clinics and shifters.

// @host      localhost:8080
// @BasePath  /

func main() {
	if !utils.LoadEnvironment() {
		slog.Default().Warn("Warning: .env file not found in common locations, using system environment variables")
	}

	log := logger.New(logger.Config{
		Format: "json",
	})

	db := database.GetPostgresConnection()
	if err := database.RunMigrations(db); err != nil {
		log.Error("Failed to run migrations", "err", err)
	}

	port := utils.GetAPIPort()

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Error("Failed to start server", "err", err)
	}
}
