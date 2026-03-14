package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
)

func NewLoginCompanyOwnerFactory() *usecases.LoginCompanyOwnerUseCase {
	companyRepository := companiesrepos.NewSqlcCompanyRepository()

	return usecases.NewLoginCompanyOwnerUseCase(companyRepository)
}
