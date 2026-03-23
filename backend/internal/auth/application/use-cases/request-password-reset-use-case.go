package usecases

import (
	"log/slog"
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
	foundVet, err := u.shiftVeterinaryRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	foundOwner, err := u.companyRepository.FindCompanyOwnerByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	explicitVet := input.UserType.Equals(sharedvalueobjects.ShiftVeterinary())
	explicitOwner := input.UserType.Equals(sharedvalueobjects.CompanyOwner())

	if explicitVet && foundVet == nil {
		return nil, &customerror.NotFoundError{Key: "Veterinary", Value: input.Email.GetValue()}
	}

	if explicitOwner && foundOwner == nil {
		return nil, &customerror.NotFoundError{Key: "Company owner", Value: input.Email.GetValue()}
	}

	var userType *sharedvalueobjects.UserType
	if foundVet != nil {
		userType = sharedvalueobjects.ShiftVeterinary()
	} else if foundOwner != nil {
		userType = sharedvalueobjects.CompanyOwner()
	}

	if userType == nil {
		return &RequestPasswordResetUseCaseOutput{Accepted: true}, nil
	}

	expiresAt := time.Now().Add(utils.GetPasswordResetTokenExpiry())
	token := uuid.Must(uuid.NewV7()).String()

	_, err = u.authRepository.CreatePasswordResetToken(
		token,
		input.Email,
		*userType,
		expiresAt,
	)
	if err != nil {
		return nil, err
	}

	resetLink := utils.GetEmailSenderBaseURL() + "/reset-password?token=" + token
	if err := u.emailSender.SendPasswordResetEmail(input.Email, resetLink); err != nil {
		slog.Error("request password reset: send email failed", "err", err)
	}

	return &RequestPasswordResetUseCaseOutput{Accepted: true}, nil
}
