package factories

import (
	companyrepos "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
	shiftvetrepos "rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/repositories"
	"rodrigoorlandini/vet-shifter/internal/auth/application/use_cases"
)

func NewLoginFactory() *usecases.LoginUseCase {
	return usecases.NewLoginUseCase(
		companyrepos.NewSqlcCompanyRepository(),
		shiftvetrepos.NewSqlcShiftVetRepository(),
	)
}
