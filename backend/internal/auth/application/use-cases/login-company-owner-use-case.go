package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
)

type LoginCompanyOwnerUseCase struct {
	companyRepository companiesrepos.CompanyRepository
}

type LoginCompanyOwnerUseCaseInput struct {
	Email      sharedvalueobjects.Email
	Password   string
	RememberMe bool
}

type LoginCompanyOwnerUseCaseOutput struct {
	AccessToken string
	ExpiresAt   string
}

func NewLoginCompanyOwnerUseCase(companyRepository companiesrepos.CompanyRepository) *LoginCompanyOwnerUseCase {
	return &LoginCompanyOwnerUseCase{
		companyRepository: companyRepository,
	}
}

func (u *LoginCompanyOwnerUseCase) Execute(input *LoginCompanyOwnerUseCaseInput) (*LoginCompanyOwnerUseCaseOutput, error) {
	if len(input.Password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidCredentialsError{}
	}

	owner, err := u.companyRepository.FindCompanyOwnerByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if owner == nil {
		return nil, &customerror.InvalidCredentialsError{}
	}

	if !utils.Argon2Compare(input.Password, owner.Password) {
		return nil, &customerror.InvalidCredentialsError{}
	}

	token, exp, err := utils.IssueJWT(
		owner.Id,
		owner.Email.GetValue(),
		sharedvalueobjects.CompanyOwner().GetValue(),
		input.RememberMe,
	)

	if err != nil {
		return nil, err
	}

	return &LoginCompanyOwnerUseCaseOutput{
		AccessToken: token,
		ExpiresAt:   exp.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
