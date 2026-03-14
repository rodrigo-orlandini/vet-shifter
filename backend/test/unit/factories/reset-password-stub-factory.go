package factories

import (
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func NewResetPasswordStubFactory() (
	*usecases.ResetPasswordUseCase,
	*repositories.StubAuthRepository,
	*repositories.StubCompanyRepository,
	*repositories.StubShiftVeterinaryRepository,
) {
	authRepo := repositories.NewStubAuthRepository()
	companyRepo := repositories.NewStubCompanyRepository()
	shiftVetRepo := repositories.NewStubShiftVeterinaryRepository()

	useCase := usecases.NewResetPasswordUseCase(
		authRepo,
		companyRepo,
		shiftVetRepo,
	)

	return useCase, authRepo, companyRepo, shiftVetRepo
}
