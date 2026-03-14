package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/repositories"
)

func NewLoginVeterinaryFactory() *usecases.LoginVeterinaryUseCase {
	shiftVetRepository := veterinariesrepos.NewSqlcShiftVeterinaryRepository()

	return usecases.NewLoginVeterinaryUseCase(shiftVetRepository)
}
