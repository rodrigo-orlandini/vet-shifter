package valueobjects_test

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueObjectCnpj(t *testing.T) {
	t.Run("it should be able to create new cnpj", func(t *testing.T) {
		cnpj, err := valueobjects.NewCnpj("00000000000100")
		assert.Nil(t, err)

		assert.Equal(t, cnpj.GetValue(), "00000000000100")
		assert.Equal(t, cnpj.GetMasked(), "00.000.000/0001-00")
	})

	t.Run("it should fail due to a small cnpj", func(t *testing.T) {
		cnpj, err := valueobjects.NewCnpj("0")
		assert.Nil(t, cnpj)
		assert.NotNil(t, err)

		assert.Equal(t, err.Error(), "Invalid value object 'Cnpj' creation with value: 0")
	})

	t.Run("it should fail due to a big cnpj", func(t *testing.T) {
		cnpj, err := valueobjects.NewCnpj("000000000000000")
		assert.Nil(t, cnpj)
		assert.NotNil(t, err)

		assert.Equal(t, err.Error(), "Invalid value object 'Cnpj' creation with value: 000000000000000")
	})
}
