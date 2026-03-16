package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func NewGetUserTypeByEmailStubFactory() (
	*usecases.GetUserTypeByEmailUseCase,
	*repositories.StubCompanyRepository,
	*repositories.StubShiftVeterinaryRepository,
) {
	companyRepo := repositories.NewStubCompanyRepository()
	shiftVetRepo := repositories.NewStubShiftVeterinaryRepository()

	return usecases.NewGetUserTypeByEmailUseCase(companyRepo, shiftVetRepo),
		companyRepo,
		shiftVetRepo
}
