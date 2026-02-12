package usecases

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/internal/companies/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
)

type RegisterCompanyUseCase struct {
	companyRepository repositories.CompanyRepository
}

type RegisterCompanyUseCaseInput struct {
	Company entities.Company
	Owner   entities.Owner
}

type RegisterCompanyUseCaseOutput struct {
	CompanyId string
}

func NewRegisterCompanyUseCase(
	companyRepository repositories.CompanyRepository,
) *RegisterCompanyUseCase {
	return &RegisterCompanyUseCase{
		companyRepository: companyRepository,
	}
}

func (u *RegisterCompanyUseCase) Execute(
	input *RegisterCompanyUseCaseInput,
) (*RegisterCompanyUseCaseOutput, error) {
	companyWithSameCnpj, err := u.companyRepository.FindByCnpj(input.Company.Cnpj)

	if err != nil {
		return nil, err
	}

	if companyWithSameCnpj != nil {
		return nil, &customerror.AlreadyExistsError{
			Entity: "Company",
			Field:  "Cnpj",
			Value:  input.Company.Cnpj.GetMasked(),
		}
	}

	ownerWithSameEmail, err := u.companyRepository.FindOwnerByEmail(input.Owner.Email)

	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Owner",
			Field:  "Email",
			Err:    err,
		}
	}

	if ownerWithSameEmail != nil {
		return nil, &customerror.AlreadyExistsError{
			Entity: "Owner",
			Field:  "Email",
			Value:  input.Owner.Email.GetValue(),
		}
	}

	company, err := u.companyRepository.Create(input.Company)

	if err != nil {
		return nil, err
	}

	passwordHash := utils.Argon2Hash(input.Owner.Password)
	input.Owner.Password = passwordHash
	input.Owner.CompanyId = company.Id

	err = u.companyRepository.RegisterOwner(input.Owner)
	if err != nil {
		return nil, err
	}
	return &RegisterCompanyUseCaseOutput{CompanyId: company.Id}, nil
}
