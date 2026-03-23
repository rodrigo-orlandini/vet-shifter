package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/veterinaries/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

type RegisterShiftVeterinaryUseCaseDependencies struct {
	ShiftVeterinaryRepository *repositories.StubShiftVeterinaryRepository
}

func NewRegisterShiftVeterinaryStubFactory() (*usecases.RegisterShiftVeterinaryUseCase, *RegisterShiftVeterinaryUseCaseDependencies) {
	shiftVeterinaryRepository := repositories.NewStubShiftVeterinaryRepository()
	useCase := usecases.NewRegisterShiftVeterinaryUseCase(shiftVeterinaryRepository)

	dependencies := &RegisterShiftVeterinaryUseCaseDependencies{
		ShiftVeterinaryRepository: shiftVeterinaryRepository,
	}

	return useCase, dependencies
}

