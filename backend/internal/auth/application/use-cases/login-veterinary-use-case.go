package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/application/repositories"
)

type LoginVeterinaryUseCase struct {
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository
}

type LoginVeterinaryUseCaseInput struct {
	Email      sharedvalueobjects.Email
	Password   string
	RememberMe bool
}

type LoginVeterinaryUseCaseOutput struct {
	AccessToken string
	ExpiresAt   string
}

func NewLoginVeterinaryUseCase(shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository) *LoginVeterinaryUseCase {
	return &LoginVeterinaryUseCase{
		shiftVeterinaryRepository: shiftVeterinaryRepository,
	}
}

func (u *LoginVeterinaryUseCase) Execute(input *LoginVeterinaryUseCaseInput) (*LoginVeterinaryUseCaseOutput, error) {
	if len(input.Password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidCredentialsError{}
	}

	veterinary, err := u.shiftVeterinaryRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if veterinary == nil {
		return nil, &customerror.InvalidCredentialsError{}
	}

	if !utils.Argon2Compare(input.Password, veterinary.Password) {
		return nil, &customerror.InvalidCredentialsError{}
	}

	token, exp, err := utils.IssueJWT(
		veterinary.Id,
		veterinary.Email.GetValue(),
		sharedvalueobjects.ShiftVeterinary().GetValue(),
		input.RememberMe,
	)

	if err != nil {
		return nil, err
	}

	return &LoginVeterinaryUseCaseOutput{
		AccessToken: token,
		ExpiresAt:   exp.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
