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
	Address      *entities.Address
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
	if input.CompanyOwner.ConsentLgpdAt == nil {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Consentimento LGPD",
			Value: "",
		}
	}

	if len(input.CompanyOwner.Password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Senha",
			Value: "deve ter ao menos 8 caracteres",
		}
	}

	companyWithSameCnpj, err := u.companyRepository.FindByCnpj(input.Company.Cnpj)

	if err != nil {
		return nil, err
	}

	if companyWithSameCnpj != nil {
		return nil, &customerror.AlreadyExistsError{
			Field: "CNPJ",
			Value: input.Company.Cnpj.GetMasked(),
		}
	}

	existingOwner, err := u.companyRepository.FindCompanyOwnerByEmail(input.CompanyOwner.Email)
	if err != nil {
		return nil, &customerror.RepositoryError{
			Entity: "Proprietário da clínica",
			Field:  "E-mail",
			Err:    err,
		}
	}

	if existingOwner != nil {
		return nil, &customerror.AlreadyExistsError{
			Field: "E-mail",
			Value: input.CompanyOwner.Email.GetValue(),
		}
	}

	var company *entities.Company

	err = u.companyRepository.InTransaction(func(repo repositories.CompanyRepository) error {
		company, err = repo.Create(input.Company)
		if err != nil {
			return err
		}

		if input.Address != nil {
			input.Address.CompanyId = company.Id

			_, err = repo.CreateAddress(*input.Address)
			if err != nil {
				return err
			}
		}

		passwordHash := utils.Argon2Hash(input.CompanyOwner.Password)
		input.CompanyOwner.Password = passwordHash
		input.CompanyOwner.CompanyId = company.Id

		return repo.RegisterCompanyOwner(input.CompanyOwner)
	})

	if err != nil {
		return nil, err
	}

	return &RegisterCompanyUseCaseOutput{CompanyId: company.Id}, nil
}
