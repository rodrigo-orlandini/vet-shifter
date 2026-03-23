package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
)

func TestEntityAddress(t *testing.T) {
	t.Run("it should be able to create new address", func(t *testing.T) {
		addr, err := entities.NewAddress("company-id-123", "Rua A", "100", "São Paulo", "SP", "01310100")
		assert.Nil(t, err)
		assert.NotNil(t, addr)
		assert.NotEmpty(t, addr.Id)
		assert.Equal(t, "company-id-123", addr.CompanyId)
		assert.Equal(t, "Rua A", addr.Street)
		assert.Equal(t, "100", addr.Number)
		assert.Equal(t, "São Paulo", addr.City)
		assert.Equal(t, "SP", addr.State.GetValue())
		assert.Equal(t, "01310100", addr.ZipCode.GetValue())
		assert.Equal(t, "01310-100", addr.ZipCode.GetMasked())
		assert.NotNil(t, addr.CreatedAt)
	})

	t.Run("it should fail due to invalid UF", func(t *testing.T) {
		addr, err := entities.NewAddress("company-id-456", "Rua B", "200", "Curitiba", "XX", "80010100")
		assert.Nil(t, addr)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "UF")
	})

	t.Run("it should fail due to invalid CEP", func(t *testing.T) {
		addr, err := entities.NewAddress("company-id-789", "Rua C", "", "Curitiba", "PR", "123")
		assert.Nil(t, addr)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "CEP")
	})

	t.Run("it should fail due to empty UF", func(t *testing.T) {
		addr, err := entities.NewAddress("company-id-000", "Rua D", "", "Curitiba", "", "80010100")
		assert.Nil(t, addr)
		assert.NotNil(t, err)
	})
}
