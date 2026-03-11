package valueobjects_test

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueObjectCpf(t *testing.T) {
	t.Run("it should be able to create new cpf", func(t *testing.T) {
		cpf, err := valueobjects.NewCpf("12345678901")
		assert.Nil(t, err)

		assert.Equal(t, "12345678901", cpf.GetValue())
		assert.Equal(t, "123.456.789-01", cpf.GetMasked())
	})

	t.Run("it should fail due to a small cpf", func(t *testing.T) {
		cpf, err := valueobjects.NewCpf("1")
		assert.Nil(t, cpf)
		assert.NotNil(t, err)

		assert.Equal(t, "Invalid value object 'Cpf' creation with value: 1", err.Error())
	})

	t.Run("it should fail due to a big cpf", func(t *testing.T) {
		cpf, err := valueobjects.NewCpf("123456789012")
		assert.Nil(t, cpf)
		assert.NotNil(t, err)

		assert.Equal(t, "Invalid value object 'Cpf' creation with value: 123456789012", err.Error())
	})
}
