package usecases_test

import (
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

func TestGetUserTypeByEmailUseCase(t *testing.T) {
	t.Run("it should return company_owner when email is owner", func(t *testing.T) {
		useCase, companyRepo, _ := factories.NewGetUserTypeByEmailStubFactory()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co")
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("00000000000")
		consent := time.Now()
		owner, _ := entities.NewCompanyOwner(*email, *phone, utils.Argon2Hash("pass"), company.Id, &consent)
		_ = companyRepo.RegisterCompanyOwner(*owner)

		out, err := useCase.Execute(&usecases.GetUserTypeByEmailUseCaseInput{Email: *email})
		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.UserType.Equals(sharedvalueobjects.CompanyOwner()))
	})

	t.Run("it should return shift_veterinary when email is veterinary", func(t *testing.T) {
		useCase, _, vetRepo := factories.NewGetUserTypeByEmailStubFactory()
		vetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "password-hash-min-8")
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")

		out, err := useCase.Execute(&usecases.GetUserTypeByEmailUseCaseInput{Email: *email})
		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.UserType.Equals(sharedvalueobjects.ShiftVeterinary()))
	})

	t.Run("it should return NotFoundError when email exists in neither", func(t *testing.T) {
		useCase, _, _ := factories.NewGetUserTypeByEmailStubFactory()
		email, _ := sharedvalueobjects.NewEmail("nobody@test.com")

		out, err := useCase.Execute(&usecases.GetUserTypeByEmailUseCaseInput{Email: *email})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.NotFoundError{}, err)
	})
}
