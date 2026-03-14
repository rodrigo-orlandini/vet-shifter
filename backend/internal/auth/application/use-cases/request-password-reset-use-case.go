package usecases

import (
	"time"

	"github.com/google/uuid"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/auth/application/services"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/application/repositories"
)

type RequestPasswordResetUseCase struct {
	authRepository            repositories.AuthRepository
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository
	companyRepository         companiesrepos.CompanyRepository
	emailSender               services.EmailSender
}

type RequestPasswordResetUseCaseInput struct {
	Email    sharedvalueobjects.Email
	UserType sharedvalueobjects.UserType
}

type RequestPasswordResetUseCaseOutput struct {
	Accepted bool
}

func NewRequestPasswordResetUseCase(
	authRepository repositories.AuthRepository,
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository,
	companyRepository companiesrepos.CompanyRepository,
	emailSender services.EmailSender,
) *RequestPasswordResetUseCase {
	return &RequestPasswordResetUseCase{
		authRepository:            authRepository,
		shiftVeterinaryRepository: shiftVeterinaryRepository,
		companyRepository:         companyRepository,
		emailSender:               emailSender,
	}
}

func (u *RequestPasswordResetUseCase) Execute(input *RequestPasswordResetUseCaseInput) (*RequestPasswordResetUseCaseOutput, error) {
	if input.UserType.Equals(sharedvalueobjects.ShiftVeterinary()) {
		foundVeterinary, err := u.shiftVeterinaryRepository.FindByEmail(input.Email)
		if err != nil {
			return nil, err
		}

		if foundVeterinary == nil {
			return nil, &customerror.NotFoundError{
				Key:   "Veterinary",
				Value: input.Email.GetValue(),
			}
		}
	} else if input.UserType.Equals(sharedvalueobjects.CompanyOwner()) {
		foundCompanyOwner, err := u.companyRepository.FindCompanyOwnerByEmail(input.Email)
		if err != nil {
			return nil, err
		}

		if foundCompanyOwner == nil {
			return nil, &customerror.NotFoundError{
				Key:   "Company owner",
				Value: input.Email.GetValue(),
			}
		}
	}

	expiresAt := time.Now().Add(utils.GetPasswordResetTokenExpiry())
	token := uuid.Must(uuid.NewV7()).String()

	_, err := u.authRepository.CreatePasswordResetToken(
		token,
		input.Email,
		input.UserType,
		expiresAt,
	)

	if err != nil {
		return nil, err
	}

	resetLink := utils.GetEmailSenderBaseURL() + "/reset-password?token=" + token

	err = u.emailSender.SendPasswordResetEmail(input.Email, resetLink)
	if err != nil {
		return nil, err
	}

	return &RequestPasswordResetUseCaseOutput{Accepted: true}, nil
}
