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
	Company      entities.Company
	CompanyOwner entities.CompanyOwner
}

type RegisterCompanyUseCaseOutput struct {
	CompanyId string
}

func NewRegisterCompanyUseCase(companyRepository repositories.CompanyRepository) *RegisterCompanyUseCase {
	return &RegisterCompanyUseCase{
		companyRepository: companyRepository,
	}
}

func (u *RegisterCompanyUseCase) Execute(
	input *RegisterCompanyUseCaseInput,
) (*RegisterCompanyUseCaseOutput, error) {
	if len(input.CompanyOwner.Password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Password",
			Value: "must be at least 8 characters",
		}
	}

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

	existingOwner, err := u.companyRepository.FindCompanyOwnerByEmail(input.CompanyOwner.Email)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "CompanyOwner",
			Field:  "Email",
			Err:    err,
		}
	}

	if existingOwner != nil {
		return nil, &customerror.AlreadyExistsError{
			Entity: "CompanyOwner",
			Field:  "Email",
			Value:  input.CompanyOwner.Email.GetValue(),
		}
	}

	company, err := u.companyRepository.Create(input.Company)

	if err != nil {
		return nil, err
	}

	passwordHash := utils.Argon2Hash(input.CompanyOwner.Password)
	input.CompanyOwner.Password = passwordHash
	input.CompanyOwner.CompanyId = company.Id

	err = u.companyRepository.RegisterCompanyOwner(input.CompanyOwner)
	if err != nil {
		return nil, err
	}

	return &RegisterCompanyUseCaseOutput{CompanyId: company.Id}, nil
}
