package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	companiesvalueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"
)

func TestRequestPasswordResetUseCase(t *testing.T) {
	t.Run("it should return Accepted when veterinary exists and send email", func(t *testing.T) {
		useCase, _, _, shiftVetRepo, emailSender := factories.NewRequestPasswordResetStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "password-hash")

		out, err := useCase.Execute(&usecases.RequestPasswordResetUseCaseInput{
			Email:    *email,
			UserType: *sharedvalueobjects.ShiftVeterinary(),
		})

		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.Accepted)
		assert.Len(t, emailSender.SentEmails, 1)
		assert.Equal(t, "vet@test.com", emailSender.SentEmails[0].To.GetValue())
		assert.Contains(t, emailSender.SentEmails[0].ResetLink, "/reset-password?token=")
	})

	t.Run("it should return Accepted when company owner exists and send email", func(t *testing.T) {
		useCase, _, companyRepo, _, emailSender := factories.NewRequestPasswordResetStubFactory()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co", nil)
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("11999999999")
		owner, _ := entities.NewCompanyOwner(*email, *phone, "password-hash", company.Id, nil)
		_ = companyRepo.RegisterCompanyOwner(*owner)

		out, err := useCase.Execute(&usecases.RequestPasswordResetUseCaseInput{
			Email:    *email,
			UserType: *sharedvalueobjects.CompanyOwner(),
		})

		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.Accepted)
		assert.Len(t, emailSender.SentEmails, 1)
		assert.Equal(t, "owner@test.com", emailSender.SentEmails[0].To.GetValue())
	})

	t.Run("it should return NotFoundError when veterinary email not found", func(t *testing.T) {
		useCase, _, _, _, _ := factories.NewRequestPasswordResetStubFactory()
		email, _ := sharedvalueobjects.NewEmail("unknown-vet@test.com")

		out, err := useCase.Execute(&usecases.RequestPasswordResetUseCaseInput{
			Email:    *email,
			UserType: *sharedvalueobjects.ShiftVeterinary(),
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.NotFoundError{}, err)
		assert.Contains(t, err.Error(), "Veterinary")
	})

	t.Run("it should return NotFoundError when company owner email not found", func(t *testing.T) {
		useCase, _, _, _, _ := factories.NewRequestPasswordResetStubFactory()
		email, _ := sharedvalueobjects.NewEmail("unknown-owner@test.com")

		out, err := useCase.Execute(&usecases.RequestPasswordResetUseCaseInput{
			Email:    *email,
			UserType: *sharedvalueobjects.CompanyOwner(),
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.NotFoundError{}, err)
		assert.Contains(t, err.Error(), "Company owner")
	})
}
