package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/repositories"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/repositories"
)

func NewResetPasswordFactory() *usecases.ResetPasswordUseCase {
	companyRepository := companiesrepos.NewSqlcCompanyRepository()
	shiftVetRepository := veterinariesrepos.NewSqlcShiftVeterinaryRepository()
	authRepository := repositories.NewSqlcAuthRepository()

	return usecases.NewResetPasswordUseCase(
		authRepository,
		companyRepository,
		shiftVetRepository,
	)
}
