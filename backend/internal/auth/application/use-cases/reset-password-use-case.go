package usecases

import (
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	autherrors "rodrigoorlandini/vet-shifter/internal/auth/application/custom-error"
	"rodrigoorlandini/vet-shifter/internal/auth/application/repositories"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	companiesentities "rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	veterinariesentities "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/application/repositories"
)

type ResetPasswordUseCase struct {
	authRepository            repositories.AuthRepository
	companyRepository         companiesrepos.CompanyRepository
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository
}

type ResetPasswordUseCaseInput struct {
	Token       string
	NewPassword string
}

type ResetPasswordUseCaseOutput struct {
	Success bool
}

func NewResetPasswordUseCase(
	authRepository repositories.AuthRepository,
	companyRepository companiesrepos.CompanyRepository,
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository,
) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		authRepository:            authRepository,
		companyRepository:         companyRepository,
		shiftVeterinaryRepository: shiftVeterinaryRepository,
	}
}

func (u *ResetPasswordUseCase) Execute(input *ResetPasswordUseCaseInput) (*ResetPasswordUseCaseOutput, error) {
	if len(input.NewPassword) < utils.MinPasswordLength {
		return nil, &customerror.InvalidCredentialsError{}
	}

	record, err := u.authRepository.GetPasswordResetToken(input.Token)
	if err != nil {
		return nil, &autherrors.InvalidResetTokenError{}
	}

	if record == nil {
		return nil, &autherrors.InvalidResetTokenError{}
	}

	if record.UsedAt != nil {
		return nil, &autherrors.InvalidResetTokenError{}
	}

	if time.Now().After(record.ExpiresAt) {
		return nil, &autherrors.InvalidResetTokenError{}
	}

	hashedPassword := utils.Argon2Hash(input.NewPassword)
	var veterinary *veterinariesentities.ShiftVeterinary
	var companyOwner *companiesentities.CompanyOwner

	if record.UserType.Equals(sharedvalueobjects.CompanyOwner()) {
		companyOwner, err = u.companyRepository.FindCompanyOwnerByEmail(record.Email)
		if err != nil {
			return nil, err
		}

		if companyOwner == nil {
			return nil, &autherrors.InvalidResetTokenError{}
		}

		err = u.companyRepository.UpdateCompanyOwnerPassword(companyOwner.Id, hashedPassword)
	} else if record.UserType.Equals(sharedvalueobjects.ShiftVeterinary()) {
		veterinary, err = u.shiftVeterinaryRepository.FindByEmail(record.Email)
		if err != nil {
			return nil, err
		}

		if veterinary == nil {
			return nil, &autherrors.InvalidResetTokenError{}
		}

		err = u.shiftVeterinaryRepository.UpdatePassword(veterinary.Id, hashedPassword)
	}

	if err != nil {
		return nil, err
	}

	err = u.authRepository.MarkPasswordResetTokenUsed(record.Id)
	if err != nil {
		return nil, err
	}

	return &ResetPasswordUseCaseOutput{Success: true}, nil
}
