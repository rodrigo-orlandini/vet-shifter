package factories

import (
	"os"

	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func NewLoginCompanyOwnerStubFactory() (
	*usecases.LoginCompanyOwnerUseCase,
	*repositories.StubCompanyRepository,
) {
	companyRepository := repositories.NewStubCompanyRepository()

	if utils.GetJWTSecret() == "" {
		_ = os.Setenv("JWT_SECRET", "test-secret-for-unit-tests")
	}

	return usecases.NewLoginCompanyOwnerUseCase(companyRepository), companyRepository
}
