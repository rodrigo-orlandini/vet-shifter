package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/repositories"
)

func NewGetUserTypeByEmailFactory() *usecases.GetUserTypeByEmailUseCase {
	companyRepository := companiesrepos.NewSqlcCompanyRepository()
	shiftVetRepository := veterinariesrepos.NewSqlcShiftVeterinaryRepository()

	return usecases.NewGetUserTypeByEmailUseCase(companyRepository, shiftVetRepository)
}
