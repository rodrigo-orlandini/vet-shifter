package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/shift_vets/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/repositories"
)

func NewRegisterShiftVetFactory() *usecases.RegisterShiftVetUseCase {
	return usecases.NewRegisterShiftVetUseCase(repositories.NewSqlcShiftVetRepository())
}
