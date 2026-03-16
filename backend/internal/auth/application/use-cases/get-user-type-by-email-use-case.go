package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	companiesrepos "rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	veterinariesrepos "rodrigoorlandini/vet-shifter/internal/veterinaries/application/repositories"
)

type GetUserTypeByEmailUseCase struct {
	companyRepository         companiesrepos.CompanyRepository
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository
}

type GetUserTypeByEmailUseCaseInput struct {
	Email sharedvalueobjects.Email
}

type GetUserTypeByEmailUseCaseOutput struct {
	UserType sharedvalueobjects.UserType
}

func NewGetUserTypeByEmailUseCase(
	companyRepository companiesrepos.CompanyRepository,
	shiftVeterinaryRepository veterinariesrepos.ShiftVeterinaryRepository,
) *GetUserTypeByEmailUseCase {
	return &GetUserTypeByEmailUseCase{
		companyRepository:         companyRepository,
		shiftVeterinaryRepository: shiftVeterinaryRepository,
	}
}

func (u *GetUserTypeByEmailUseCase) Execute(input *GetUserTypeByEmailUseCaseInput) (*GetUserTypeByEmailUseCaseOutput, error) {
	owner, err := u.companyRepository.FindCompanyOwnerByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if owner != nil {
		return &GetUserTypeByEmailUseCaseOutput{UserType: *sharedvalueobjects.CompanyOwner()}, nil
	}

	vet, err := u.shiftVeterinaryRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if vet != nil {
		return &GetUserTypeByEmailUseCaseOutput{UserType: *sharedvalueobjects.ShiftVeterinary()}, nil
	}

	return nil, &customerror.NotFoundError{Key: "User", Value: input.Email.GetValue()}
}
