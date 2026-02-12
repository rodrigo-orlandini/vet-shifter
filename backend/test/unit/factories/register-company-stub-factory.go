package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/companies/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

type RegisterCompanyUseCaseDependencies struct {
	CompanyRepository *repositories.StubCompanyRepository
}

func NewRegisterCompanyStubFactory() (*usecases.RegisterCompanyUseCase, *RegisterCompanyUseCaseDependencies) {
	companyRepository := repositories.NewStubCompanyRepository()
	useCase := usecases.NewRegisterCompanyUseCase(companyRepository)

	dependencies := &RegisterCompanyUseCaseDependencies{
		CompanyRepository: companyRepository,
	}

	return useCase, dependencies
}
