package usecases_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/test/unit/factories"
)

func TestLoginVeterinaryUseCase(t *testing.T) {
	if os.Getenv("JWT_SECRET") == "" {
		_ = os.Setenv("JWT_SECRET", "test-secret")
	}

	t.Run("it should return token when veterinary credentials are correct", func(t *testing.T) {
		useCase, stub := factories.NewLoginVeterinaryStubFactory()
		hash := utils.Argon2Hash("password123")
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		stub.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, hash)

		out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
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
		useCase, stub := factories.NewLoginVeterinaryStubFactory()
		hash := utils.Argon2Hash("correct-password")
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		stub.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, hash)

		out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
			Email:    *email,
			Password: "wrong-password",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when email is not found", func(t *testing.T) {
		useCase, _ := factories.NewLoginVeterinaryStubFactory()
		unknownEmail, _ := sharedvalueobjects.NewEmail("unknown@test.com")

		out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
			Email:    *unknownEmail,
			Password: "anypassword",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when password is shorter than 8 chars", func(t *testing.T) {
		useCase, stub := factories.NewLoginVeterinaryStubFactory()
		hash := utils.Argon2Hash("password123")
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		stub.AddUser("vet@test.com", "vet-id-1", sharedvalueobjects.UserTypeShiftVeterinary, hash)

		out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
			Email:    *email,
			Password: "short",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})

	t.Run("it should return InvalidCredentialsError when user is owner not veterinary", func(t *testing.T) {
		useCase, stub := factories.NewLoginVeterinaryStubFactory()
		hash := utils.Argon2Hash("password123")
		ownerEmail, _ := sharedvalueobjects.NewEmail("owner@test.com")
		stub.AddUser("owner@test.com", "owner-id-1", sharedvalueobjects.UserTypeCompanyOwner, hash)

		out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
			Email:    *ownerEmail,
			Password: "password123",
		})
		assert.Nil(t, out)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidCredentialsError{}, err)
	})
}
