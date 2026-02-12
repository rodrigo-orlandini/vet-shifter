package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/shifts/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shifts/infrastructure/repositories"
)

func NewCreateShiftFactory() *usecases.CreateShiftUseCase {
	return usecases.NewCreateShiftUseCase(repositories.NewSqlcShiftRepository())
}

func NewListShiftsFactory() *usecases.ListShiftsUseCase {
	return usecases.NewListShiftsUseCase(repositories.NewSqlcShiftRepository())
}
