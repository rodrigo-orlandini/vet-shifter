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
// @description     API para clínicas veterinárias e plantonistas.

// @host      localhost:8000
// @BasePath  /

func main() {
	envLoaded := utils.LoadEnvironment()

	logger.SetDefault(logger.New(logger.Config{
		Format: "json",
	}))

	if !envLoaded {
		slog.Warn("Arquivo .env não encontrado nos caminhos usuais; usando variáveis de ambiente do sistema")
	}

	db := database.GetPostgresConnection()
	if err := database.RunMigrations(db); err != nil {
		slog.Error("Falha ao executar migrations", "err", err)
	}

	port := utils.GetAPIPort()

	r := router.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		slog.Error("Falha ao iniciar o servidor", "err", err)
	}
}
