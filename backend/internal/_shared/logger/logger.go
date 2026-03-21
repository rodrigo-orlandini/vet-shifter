package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type Env string

const (
	EnvProduction  Env = "production"
	EnvStaging     Env = "staging"
	EnvDevelopment Env = "development"
)

type Config struct {
	Format string
	Output io.Writer
	Env    Env
}

func New(config Config) *slog.Logger {
	if config.Output == nil {
		config.Output = os.Stdout
	}

	env := config.Env
	if env == "" {
		env = envFromOS()
	}

	level := levelForEnv(env)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler
	switch strings.ToLower(config.Format) {
	case "json":
		handler = slog.NewJSONHandler(config.Output, opts)
	default:
		handler = slog.NewTextHandler(config.Output, opts)
	}

	return slog.New(handler)
}

func envFromOS() Env {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("APP_ENV")))
	switch v {
	case "production", "prod":
		return EnvProduction
	case "staging", "stage":
		return EnvStaging
	case "development", "dev":
		return EnvDevelopment
	default:
		return EnvDevelopment
	}
}

func levelForEnv(env Env) slog.Level {
	switch env {
	case EnvProduction:
		return slog.LevelInfo
	case EnvStaging, EnvDevelopment:
		return slog.LevelDebug
	default:
		return slog.LevelDebug
	}
}

func SetDefault(l *slog.Logger) {
	slog.SetDefault(l)
}
