package usecases_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	autherrors "rodrigoorlandini/vet-shifter/internal/auth/application/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	companiesvalueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"rodrigoorlandini/vet-shifter/test/unit/factories"
)

func TestResetPasswordUseCase(t *testing.T) {
	t.Run("it should return Success when token is valid for veterinary", func(t *testing.T) {
		useCase, authRepo, _, shiftVetRepo := factories.NewResetPasswordStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "old-hash")
		expiresAt := time.Now().Add(time.Hour)
		_, _ = authRepo.CreatePasswordResetToken("reset-token-123", *email, *sharedvalueobjects.ShiftVeterinary(), expiresAt)

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "reset-token-123",
			NewPassword: "newpassword123",
		})

		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.Success)
	})

	t.Run("it should return Success when token is valid for company owner", func(t *testing.T) {
		useCase, authRepo, companyRepo, _ := factories.NewResetPasswordStubFactory()
		cnpj, _ := companiesvalueobjects.NewCnpj("00000000000100")
		company, _ := entities.NewCompany(*cnpj, "Test Co", nil)
		companyRepo.Create(*company)
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		phone, _ := sharedvalueobjects.NewPhone("11999999999")
		owner, _ := entities.NewCompanyOwner(*email, *phone, "old-hash", company.Id, nil)
		_ = companyRepo.RegisterCompanyOwner(*owner)
		expiresAt := time.Now().Add(time.Hour)
		_, _ = authRepo.CreatePasswordResetToken("reset-token-456", *email, *sharedvalueobjects.CompanyOwner(), expiresAt)

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "reset-token-456",
			NewPassword: "newpassword123",
		})

		assert.Nil(t, err)
		assert.NotNil(t, out)
		assert.True(t, out.Success)
	})

	t.Run("it should return InvalidResetTokenError when token not found", func(t *testing.T) {
		useCase, _, _, _ := factories.NewResetPasswordStubFactory()

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "unknown-token",
			NewPassword: "newpassword123",
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &autherrors.InvalidResetTokenError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when new password is shorter than 8 chars", func(t *testing.T) {
		useCase, authRepo, _, shiftVetRepo := factories.NewResetPasswordStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "oldpassword")
		expiresAt := time.Now().Add(time.Hour)
		_, _ = authRepo.CreatePasswordResetToken("token-short-pw", *email, *sharedvalueobjects.ShiftVeterinary(), expiresAt)

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "token-short-pw",
			NewPassword: "short",
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidResetTokenError when token already used", func(t *testing.T) {
		useCase, authRepo, _, shiftVetRepo := factories.NewResetPasswordStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "oldpassword")
		expiresAt := time.Now().Add(time.Hour)
		rec, _ := authRepo.CreatePasswordResetToken("already-used-token", *email, *sharedvalueobjects.ShiftVeterinary(), expiresAt)
		_ = authRepo.MarkPasswordResetTokenUsed(rec.Id)

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "already-used-token",
			NewPassword: "newpassword123",
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &autherrors.InvalidResetTokenError{}, err)
	})

	t.Run("it should return InvalidResetTokenError when token is expired", func(t *testing.T) {
		useCase, authRepo, _, shiftVetRepo := factories.NewResetPasswordStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "oldpassword")
		expiresAt := time.Now().Add(-time.Hour)
		_, _ = authRepo.CreatePasswordResetToken("expired-token", *email, *sharedvalueobjects.ShiftVeterinary(), expiresAt)

		out, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "expired-token",
			NewPassword: "newpassword123",
		})

		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &autherrors.InvalidResetTokenError{}, err)
	})

	t.Run("it should update password and mark token used", func(t *testing.T) {
		useCase, authRepo, _, shiftVetRepo := factories.NewResetPasswordStubFactory()
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		shiftVetRepo.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, "old-hash")
		expiresAt := time.Now().Add(time.Hour)
		_, _ = authRepo.CreatePasswordResetToken("token-to-use", *email, *sharedvalueobjects.ShiftVeterinary(), expiresAt)

		_, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
			Token:       "token-to-use",
			NewPassword: "newpassword123",
		})
		assert.Nil(t, err)

		record, _ := authRepo.GetPasswordResetToken("token-to-use")
		assert.NotNil(t, record)
		assert.NotNil(t, record.UsedAt)

		vet, _ := shiftVetRepo.FindByEmail(*email)
		assert.NotNil(t, vet)
		assert.True(t, utils.Argon2Compare("newpassword123", vet.Password))
	})
}
