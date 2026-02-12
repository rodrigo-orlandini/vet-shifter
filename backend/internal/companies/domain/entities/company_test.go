package entities_test

import (
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityCompany(t *testing.T) {
	t.Run("it should be able to create new company", func(t *testing.T) {
		cnpj, err := valueobjects.NewCnpj("00000000000100")
		assert.Nil(t, err)

		company, err := entities.NewCompany(*cnpj, "Test company", nil)
		assert.Nil(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, company.Name, "Test company")
		assert.Equal(t, company.Cnpj.GetValue(), "00000000000100")
		assert.NotNil(t, company.Id)
		assert.NotNil(t, company.CreatedAt)
	})
}
