package usecases_test

import (
	"testing"
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/companies/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseRegisterCompany(t *testing.T) {
	t.Run("it should be able to register a new company", func(t *testing.T) {
		useCase, deps := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, err := entities.NewCompany(*cnpj, "Test Company")
		assert.Nil(t, err)

		email, _ := sharedvalueobjects.NewEmail("test@email.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		companyOwner, err := entities.NewCompanyOwner(*email, *phone, "password", company.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *company,
			CompanyOwner: *companyOwner,
		})
		assert.Nil(t, err)

		company, err = deps.CompanyRepository.FindByCnpj(*cnpj)
		assert.Nil(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, company.Name, "Test Company")
		assert.Equal(t, company.Cnpj.GetValue(), "00000000000100")
		assert.NotNil(t, company.Id)
		assert.NotNil(t, company.CreatedAt)

		companyOwner, err = deps.CompanyRepository.FindCompanyOwnerByEmail(*email)
		assert.Nil(t, err)
		assert.NotNil(t, companyOwner)
		assert.Equal(t, companyOwner.Email.GetValue(), "test@email.com")
		assert.Equal(t, companyOwner.Phone.GetValue(), "00000000000")
		assert.Equal(t, companyOwner.CompanyId, company.Id)
	})

	t.Run("it should be able to register a company with address", func(t *testing.T) {
		useCase, deps := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, err := entities.NewCompany(*cnpj, "Test Company")
		assert.Nil(t, err)

		email, _ := sharedvalueobjects.NewEmail("test@email.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		companyOwner, err := entities.NewCompanyOwner(*email, *phone, "password", company.Id, &consent)
		assert.Nil(t, err)

		address, err := entities.NewAddress(company.Id, "Rua A", "100", "São Paulo", "SP", "01310100")
		assert.Nil(t, err)

		out, err := useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *company,
			CompanyOwner: *companyOwner,
			Address:      address,
		})
		assert.Nil(t, err)
		assert.NotNil(t, out)

		savedAddr, err := deps.CompanyRepository.FindAddressByCompanyID(out.CompanyId)
		assert.Nil(t, err)
		assert.NotNil(t, savedAddr)
		assert.Equal(t, "Rua A", savedAddr.Street)
		assert.Equal(t, "100", savedAddr.Number)
		assert.Equal(t, "São Paulo", savedAddr.City)
		assert.Equal(t, "SP", savedAddr.State.GetValue())
		assert.Equal(t, "01310100", savedAddr.ZipCode.GetValue())
		assert.Equal(t, "01310-100", savedAddr.ZipCode.GetMasked())
	})

	t.Run("it should not be able to register a company with an existing cnpj", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		existingCompany, err := entities.NewCompany(*cnpj, "Existing Company")
		assert.Nil(t, err)

		existingEmail, _ := sharedvalueobjects.NewEmail("existing@email.com")
		existingPhone, _ := sharedvalueobjects.NewPhone("00000000000")
		existingOwner, err := entities.NewCompanyOwner(*existingEmail, *existingPhone, "password", existingCompany.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *existingCompany,
			CompanyOwner: *existingOwner,
		})
		assert.Nil(t, err)

		duplicateCompany, err := entities.NewCompany(*cnpj, "Duplicate Company")
		assert.Nil(t, err)

		duplicateEmail, _ := sharedvalueobjects.NewEmail("duplicate@email.com")
		duplicatePhone, _ := sharedvalueobjects.NewPhone("11111111111")
		duplicateOwner, err := entities.NewCompanyOwner(*duplicateEmail, *duplicatePhone, "password", duplicateCompany.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *duplicateCompany,
			CompanyOwner: *duplicateOwner,
		})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "CNPJ", alreadyExistsErr.Field)
	})

	t.Run("it should not be able to register an owner with an existing email", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj1, _ := valueobjects.NewCnpj("00000000000100")
		existingCompany, err := entities.NewCompany(*cnpj1, "Existing Company")
		assert.Nil(t, err)

		email, _ := sharedvalueobjects.NewEmail("existing@email.com")
		phone1, _ := sharedvalueobjects.NewPhone("00000000000")
		existingOwner, err := entities.NewCompanyOwner(*email, *phone1, "password", existingCompany.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *existingCompany,
			CompanyOwner: *existingOwner,
		})
		assert.Nil(t, err)

		cnpj2, _ := valueobjects.NewCnpj("00000000000200")
		duplicateCompany, err := entities.NewCompany(*cnpj2, "Duplicate Company")
		assert.Nil(t, err)

		phone2, _ := sharedvalueobjects.NewPhone("11111111111")
		duplicateOwner, err := entities.NewCompanyOwner(*email, *phone2, "password", duplicateCompany.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *duplicateCompany,
			CompanyOwner: *duplicateOwner,
		})
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.AlreadyExistsError{}, err)
		alreadyExistsErr := err.(*customerror.AlreadyExistsError)
		assert.Equal(t, "E-mail", alreadyExistsErr.Field)
		assert.Equal(t, "existing@email.com", alreadyExistsErr.Value)
	})

	t.Run("it should not be able to register without LGPD consent", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, err := entities.NewCompany(*cnpj, "Test Company")
		assert.Nil(t, err)

		email, _ := sharedvalueobjects.NewEmail("test@email.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		companyOwner, err := entities.NewCompanyOwner(*email, *phone, "password", company.Id, &consent)
		assert.Nil(t, err)

		companyOwner.ConsentLgpdAt = nil

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *company,
			CompanyOwner: *companyOwner,
		})

		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidValueObjectError{}, err)
		invalidErr := err.(*customerror.InvalidValueObjectError)
		assert.Equal(t, "Consentimento LGPD", invalidErr.Key)
	})

	t.Run("it should not be able to register when password has less than 8 characters", func(t *testing.T) {
		useCase, _ := factories.NewRegisterCompanyStubFactory()
		consent := time.Now()

		cnpj, _ := valueobjects.NewCnpj("00000000000100")
		company, err := entities.NewCompany(*cnpj, "Test Company")
		assert.Nil(t, err)

		email, _ := sharedvalueobjects.NewEmail("test@email.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		companyOwner, err := entities.NewCompanyOwner(*email, *phone, "short", company.Id, &consent)
		assert.Nil(t, err)

		_, err = useCase.Execute(&usecases.RegisterCompanyUseCaseInput{
			Company:      *company,
			CompanyOwner: *companyOwner,
		})

		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidValueObjectError{}, err)
		invalidErr := err.(*customerror.InvalidValueObjectError)
		assert.Equal(t, "Senha", invalidErr.Key)
		assert.Equal(t, "deve ter ao menos 8 caracteres", invalidErr.Value)
	})
}
