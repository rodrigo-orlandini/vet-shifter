package factories

import (
	"os"

	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func NewLoginVeterinaryStubFactory() (
	*usecases.LoginVeterinaryUseCase,
	*repositories.StubShiftVeterinaryRepository,
) {
	shiftVeterinaryRepository := repositories.NewStubShiftVeterinaryRepository()

	if utils.GetJWTSecret() == "" {
		_ = os.Setenv("JWT_SECRET", "test-secret-for-unit-tests")
	}

	return usecases.NewLoginVeterinaryUseCase(shiftVeterinaryRepository), shiftVeterinaryRepository
}
