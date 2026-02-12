package utils

import (
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvironment() bool {
	envPaths := []string{
		".env",
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
		"backend/.env",
	}

	var loaded bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			loaded = true
			break
		}
	}

	return loaded
}
