package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func NewRequestPasswordResetStubFactory() (
	*usecases.RequestPasswordResetUseCase,
	*repositories.StubAuthRepository,
	*repositories.StubCompanyRepository,
	*repositories.StubShiftVeterinaryRepository,
	*repositories.StubEmailSender,
) {
	authRepo := repositories.NewStubAuthRepository()
	companyRepo := repositories.NewStubCompanyRepository()
	shiftVetRepo := repositories.NewStubShiftVeterinaryRepository()
	emailSender := repositories.NewStubEmailSender()

	useCase := usecases.NewRequestPasswordResetUseCase(
		authRepo,
		shiftVetRepo,
		companyRepo,
		emailSender,
	)

	return useCase, authRepo, companyRepo, shiftVetRepo, emailSender
}
