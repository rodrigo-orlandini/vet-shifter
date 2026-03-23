package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestUserType(t *testing.T) {
	t.Run("it should be able to create a company owner user type", func(t *testing.T) {
		ut, err := valueobjects.NewUserType(valueobjects.UserTypeCompanyOwner)
		assert.Nil(t, err)
		assert.NotNil(t, ut)
		assert.Equal(t, valueobjects.UserTypeCompanyOwner, ut.GetValue())
		assert.Equal(t, valueobjects.UserTypeCompanyOwner, valueobjects.CompanyOwner().GetValue())

		same, _ := valueobjects.NewUserType(valueobjects.UserTypeCompanyOwner)
		assert.True(t, ut.Equals(same))
		other, _ := valueobjects.NewUserType(valueobjects.UserTypeShiftVeterinary)
		assert.False(t, ut.Equals(other))
		assert.False(t, ut.Equals(nil))
	})

	t.Run("it should be able to create a shift veterinary user type", func(t *testing.T) {
		ut, err := valueobjects.NewUserType(valueobjects.UserTypeShiftVeterinary)
		assert.Nil(t, err)
		assert.NotNil(t, ut)
		assert.Equal(t, valueobjects.UserTypeShiftVeterinary, ut.GetValue())
		assert.Equal(t, valueobjects.UserTypeShiftVeterinary, valueobjects.ShiftVeterinary().GetValue())

		same, _ := valueobjects.NewUserType(valueobjects.UserTypeShiftVeterinary)
		assert.True(t, ut.Equals(same))
		other, _ := valueobjects.NewUserType(valueobjects.UserTypeCompanyOwner)
		assert.False(t, ut.Equals(other))
		assert.False(t, ut.Equals(nil))
	})

	t.Run("it should fail for invalid user type", func(t *testing.T) {
		ut, err := valueobjects.NewUserType("invalid")
		assert.Nil(t, ut)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Tipo de Usuário")
	})

	t.Run("it should fail for empty string", func(t *testing.T) {
		ut, err := valueobjects.NewUserType("")
		assert.Nil(t, ut)
		assert.NotNil(t, err)
	})
}
