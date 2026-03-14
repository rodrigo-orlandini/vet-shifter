package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/companies/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
)

func NewRegisterCompanyFactory() *usecases.RegisterCompanyUseCase {
	companyRepository := repositories.NewSqlcCompanyRepository()

	return usecases.NewRegisterCompanyUseCase(companyRepository)
}
