package entities_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestEntityCompanyOwner(t *testing.T) {
	t.Run("it should be able to create new company owner", func(t *testing.T) {
		email, err := sharedvalueobjects.NewEmail("test@email.com")
		assert.Nil(t, err)

		companyId, err := uuid.NewV7()
		assert.Nil(t, err)

		phone, err := sharedvalueobjects.NewPhone("00000000000")
		assert.Nil(t, err)

		owner, err := entities.NewCompanyOwner(*email, *phone, "hashed", companyId.String(), nil)
		assert.Nil(t, err)
		assert.NotNil(t, owner)
		assert.NotNil(t, owner.Id)
		assert.Equal(t, owner.Email.GetValue(), "test@email.com")
		assert.Equal(t, owner.Phone.GetValue(), "00000000000")
		assert.Equal(t, owner.Password, "hashed")
		assert.Equal(t, owner.CompanyId, companyId.String())
	})
}
