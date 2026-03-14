package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
)

func TestPasswordResetToken(t *testing.T) {
	t.Run("it should create password reset token with company owner type", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		userType := sharedvalueobjects.CompanyOwner()
		expiresAt := time.Now().Add(time.Hour)

		token, err := entities.NewPasswordResetToken(
			"token-id-1",
			"reset-token-abc",
			*email,
			*userType,
			expiresAt,
			nil,
		)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, "token-id-1", token.Id)
		assert.Equal(t, "reset-token-abc", token.Token)
		assert.Equal(t, "owner@test.com", token.Email.GetValue())
		assert.True(t, token.UserType.Equals(userType))
		assert.Equal(t, expiresAt.Unix(), token.ExpiresAt.Unix())
		assert.Nil(t, token.UsedAt)
	})

	t.Run("it should create password reset token with shift veterinary type", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		userType := sharedvalueobjects.ShiftVeterinary()
		expiresAt := time.Now().Add(2 * time.Hour)

		token, err := entities.NewPasswordResetToken(
			"token-id-2",
			"reset-token-xyz",
			*email,
			*userType,
			expiresAt,
			nil,
		)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, "token-id-2", token.Id)
		assert.Equal(t, "vet@test.com", token.Email.GetValue())
		assert.True(t, token.UserType.Equals(userType))
		assert.Nil(t, token.UsedAt)
	})

	t.Run("it should create password reset token with usedAt already set", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("user@test.com")
		userType := sharedvalueobjects.CompanyOwner()
		expiresAt := time.Now().Add(time.Hour)
		usedAt := time.Now()

		token, err := entities.NewPasswordResetToken(
			"token-id-3",
			"used-token",
			*email,
			*userType,
			expiresAt,
			&usedAt,
		)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.NotNil(t, token.UsedAt)
		assert.Equal(t, usedAt.Unix(), token.UsedAt.Unix())
	})
}
