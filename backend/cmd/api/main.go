package main

import (
	"log"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
)

func main() {
	if !utils.LoadEnvironment() {
		log.Println("Warning: .env file not found in common locations, using system environment variables")
	}

	db := database.GetPostgresConnection()
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
