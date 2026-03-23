package usecases_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	companiesvalueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"
)

func TestLoginCompanyOwnerUseCase(t *testing.T) {
	if os.Getenv("JWT_SECRET") == "" {
		_ = os.Setenv("JWT_SECRET", "test-secret")
	}

	t.Run("it should return token when owner credentials are correct", func(t *testing.T) {
		useCase, companyRepo := factories.NewLoginCompanyOwnerStubFactory()
		consent := time.Now()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co")
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		owner, _ := entities.NewCompanyOwner(*email, *phone, utils.Argon2Hash("password123"), company.Id, &consent)
		_ = companyRepo.RegisterCompanyOwner(*owner)

		out, err := useCase.Execute(&usecases.LoginCompanyOwnerUseCaseInput{
			Email:      *email,
			Password:   "password123",
			RememberMe: false,
		})
		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.NotEmpty(t, out.AccessToken)
		assert.NotEmpty(t, out.ExpiresAt)
	})

	t.Run("it should return InvalidCredentialsError when password is wrong", func(t *testing.T) {
		useCase, companyRepo := factories.NewLoginCompanyOwnerStubFactory()
		consent := time.Now()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co")
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		owner, _ := entities.NewCompanyOwner(*email, *phone, utils.Argon2Hash("correct-password"), company.Id, &consent)
		_ = companyRepo.RegisterCompanyOwner(*owner)

		out, err := useCase.Execute(&usecases.LoginCompanyOwnerUseCaseInput{
			Email:    *email,
			Password: "wrong-password",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when email is not found", func(t *testing.T) {
		useCase, _ := factories.NewLoginCompanyOwnerStubFactory()
		unknownEmail, _ := sharedvalueobjects.NewEmail("unknown@test.com")

		out, err := useCase.Execute(&usecases.LoginCompanyOwnerUseCaseInput{
			Email:    *unknownEmail,
			Password: "anypassword",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when password is shorter than 8 chars", func(t *testing.T) {
		useCase, companyRepo := factories.NewLoginCompanyOwnerStubFactory()
		consent := time.Now()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co")
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		owner, _ := entities.NewCompanyOwner(*email, *phone, "hash", company.Id, &consent)
		_ = companyRepo.RegisterCompanyOwner(*owner)

		out, err := useCase.Execute(&usecases.LoginCompanyOwnerUseCaseInput{
			Email:    *email,
			Password: "short",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})
}
