package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestGenericUser(t *testing.T) {
	t.Run("it should create generic user with company owner type", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("owner@test.com")
		userType := sharedvalueobjects.CompanyOwner()
		password := "password123"

		u, err := entities.NewGenericUser("user-id-1", *email, password, *userType)
		assert.Nil(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, "user-id-1", u.Id)
		assert.Equal(t, "owner@test.com", u.Email.GetValue())
		assert.Equal(t, password, u.Password)
		assert.True(t, u.Type.Equals(userType))
	})

	t.Run("it should create generic user with shift veterinary type", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("vet@test.com")
		userType := sharedvalueobjects.ShiftVeterinary()
		password := "validpass12"

		u, err := entities.NewGenericUser("vet-id-1", *email, password, *userType)
		assert.Nil(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, "vet-id-1", u.Id)
		assert.Equal(t, "vet@test.com", u.Email.GetValue())
		assert.Equal(t, password, u.Password)
		assert.True(t, u.Type.Equals(userType))
	})

	t.Run("it should fail for password shorter than min length", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("user@test.com")
		userType := sharedvalueobjects.CompanyOwner()

		u, err := entities.NewGenericUser("user-id-1", *email, "short", *userType)
		assert.Nil(t, u)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Invalid value object")
		assert.Contains(t, err.Error(), "Password")
	})
}
