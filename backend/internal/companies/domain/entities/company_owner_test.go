package entities_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
)

func TestEntityCompanyOwner(t *testing.T) {
	t.Run("it should be able to create new company owner", func(t *testing.T) {
		email, err := sharedvalueobjects.NewEmail("test@email.com")
		assert.Nil(t, err)

		companyId, err := uuid.NewV7()
		assert.Nil(t, err)

		phone, err := sharedvalueobjects.NewPhone("00000000000")
		assert.Nil(t, err)

		consent := time.Now()
		owner, err := entities.NewCompanyOwner(*email, *phone, "hashed", companyId.String(), &consent)
		assert.Nil(t, err)
		assert.NotNil(t, owner)
		assert.NotNil(t, owner.Id)
		assert.Equal(t, owner.Email.GetValue(), "test@email.com")
		assert.Equal(t, owner.Phone.GetValue(), "00000000000")
		assert.Equal(t, owner.Password, "hashed")
		assert.Equal(t, owner.CompanyId, companyId.String())
		assert.NotNil(t, owner.ConsentLgpdAt)
	})

	t.Run("it should fail without LGPD consent", func(t *testing.T) {
		email, _ := sharedvalueobjects.NewEmail("test@email.com")
		companyId, _ := uuid.NewV7()
		phone, _ := sharedvalueobjects.NewPhone("00000000000")

		owner, err := entities.NewCompanyOwner(*email, *phone, "hashed", companyId.String(), nil)
		assert.Nil(t, owner)
		assert.NotNil(t, err)
		assert.IsType(t, &customerror.InvalidValueObjectError{}, err)
		assert.Equal(t, "Consentimento LGPD", err.(*customerror.InvalidValueObjectError).Key)
	})
}
