package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/veterinaries/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/repositories"
)

func NewRegisterShiftVeterinaryFactory() *usecases.RegisterShiftVeterinaryUseCase {
	shiftVeterinaryRepository := repositories.NewSqlcShiftVeterinaryRepository()

	return usecases.NewRegisterShiftVeterinaryUseCase(shiftVeterinaryRepository)
}

