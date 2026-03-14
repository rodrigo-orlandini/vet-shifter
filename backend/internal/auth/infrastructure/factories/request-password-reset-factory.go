package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/email"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/repositories"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/repositories"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/repositories"
)

func NewRequestPasswordResetFactory() *usecases.RequestPasswordResetUseCase {
	authRepository := repositories.NewSqlcAuthRepository()
	shiftVetRepository := veterinariesrepos.NewSqlcShiftVeterinaryRepository()
	companyRepository := companiesrepos.NewSqlcCompanyRepository()

	return usecases.NewRequestPasswordResetUseCase(
		authRepository,
		shiftVetRepository,
		companyRepository,
		email.NewResendSender(),
	)
}
