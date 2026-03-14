package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnvironment() bool {
	return loadEnvFile([]string{
		".env",
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
		"backend/.env",
	})
}

func LoadTestEnvironment() bool {
	return loadEnvFile([]string{
		".env.test",
		filepath.Join("..", ".env.test"),
		filepath.Join("..", "..", ".env.test"),
		"backend/.env.test",
	})
}

func loadEnvFile(paths []string) bool {
	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			return true
		}
	}

	return false
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetAPIPort() string {
	if p := os.Getenv("API_PORT"); p != "" {
		return p
	}
	return "8080"
}

func GetEmailSenderBaseURL() string {
	if u := os.Getenv("EMAIL_SENDER_BASE_URL"); u != "" {
		return u
	}

	return "http://localhost:3000"
}

func GetPasswordResetTokenExpiry() time.Duration {
	s := os.Getenv("PASSWORD_RESET_TOKEN_EXPIRY_HOURS")
	if s == "" {
		return 1 * time.Hour
	}

	h, err := strconv.Atoi(s)
	if err != nil || h <= 0 {
		return 1 * time.Hour
	}

	return time.Duration(h) * time.Hour
}

func GetEmailSenderAPIKey() string {
	return os.Getenv("EMAIL_SENDER_API_KEY")
}

func GetEmailSenderFromEmail() string {
	if from := os.Getenv("EMAIL_SENDER_FROM_EMAIL"); from != "" {
		return from
	}

	return "Vet Shifter <onboarding@resend.dev>"
}
