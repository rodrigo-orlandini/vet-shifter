package usecases_test

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/companies/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseRegisterCompany(t *testing.T) {
	t.Run("it should be able to register a new company", func(t *testing.T) {
		useCase, deps := factories.NewRegisterCompanyStubFactory()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, err := entities.NewCompany(*cnpj, "Test Company", nil)
		assert.Nil(t, err)

		email, _ := valueobjects.NewEmail("test@email.com")
		phone, _ := valueobjects.NewPhone("00000000000")
		owner, err := entities.NewOwner(*email, *phone, "hashed", company.Id, nil)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company: *company,
			Owner:   *owner,
		})
		assert.Nil(t, err)

		company, err = deps.CompanyRepository.FindByCnpj(*cnpj)
		assert.Nil(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, company.Name, "Test Company")
		assert.Equal(t, company.Cnpj.GetValue(), "00000000000100")
		assert.NotNil(t, company.Id)
		assert.NotNil(t, company.CreatedAt)

		owner, err = deps.CompanyRepository.FindOwnerByEmail(*email)
		assert.Nil(t, err)
		assert.NotNil(t, owner)
		assert.Equal(t, owner.Email.GetValue(), "test@email.com")
		assert.Equal(t, owner.Phone.GetValue(), "00000000000")
		assert.Equal(t, owner.CompanyId, company.Id)
	})

	t.Run("it should not be able to register a company with an existing cnpj", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		existingCompany, err := entities.NewCompany(*cnpj, "Existing Company", nil)
		assert.Nil(t, err)

		existingEmail, _ := valueobjects.NewEmail("existing@email.com")
		existingPhone, _ := valueobjects.NewPhone("00000000000")
		existingOwner, err := entities.NewOwner(*existingEmail, *existingPhone, "hashed", existingCompany.Id, nil)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company: *existingCompany,
			Owner:   *existingOwner,
		})
		assert.Nil(t, err)

		duplicateCompany, err := entities.NewCompany(*cnpj, "Duplicate Company", nil)
		assert.Nil(t, err)

		duplicateEmail, _ := valueobjects.NewEmail("duplicate@email.com")
		duplicatePhone, _ := valueobjects.NewPhone("11111111111")
		duplicateOwner, err := entities.NewOwner(*duplicateEmail, *duplicatePhone, "hashed", duplicateCompany.Id, nil)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company: *duplicateCompany,
			Owner:   *duplicateOwner,
		})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "Company", alreadyExistsErr.Entity)
		assert.Equal(t, "Cnpj", alreadyExistsErr.Field)
	})

	t.Run("it should not be able to register an owner with an existing email", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()

		cnpj1, _ := valueobjects.NewCnpj("00000000000100")
		existingCompany, err := entities.NewCompany(*cnpj1, "Existing Company", nil)
		assert.Nil(t, err)

		email, _ := valueobjects.NewEmail("existing@email.com")
		phone1, _ := valueobjects.NewPhone("00000000000")
		existingOwner, err := entities.NewOwner(*email, *phone1, "hashed", existingCompany.Id, nil)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company: *existingCompany,
			Owner:   *existingOwner,
		})
		assert.Nil(t, err)

		cnpj2, _ := valueobjects.NewCnpj("00000000000200")
		duplicateCompany, err := entities.NewCompany(*cnpj2, "Duplicate Company", nil)
		assert.Nil(t, err)

		phone2, _ := valueobjects.NewPhone("11111111111")
		duplicateOwner, err := entities.NewOwner(*email, *phone2, "hashed", duplicateCompany.Id, nil)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company: *duplicateCompany,
			Owner:   *duplicateOwner,
		})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "Owner", alreadyExistsErr.Entity)
		assert.Equal(t, "Email", alreadyExistsErr.Field)
		assert.Equal(t, "existing@email.com", alreadyExistsErr.Value)
	})
}
