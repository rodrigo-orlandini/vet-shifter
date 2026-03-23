package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestValueObjectCep(t *testing.T) {
	t.Run("it should be able to create new cep", func(t *testing.T) {
		cep, err := valueobjects.NewCep("01310100")
		assert.Nil(t, err)

		assert.Equal(t, "01310100", cep.GetValue())
		assert.Equal(t, "01310-100", cep.GetMasked())
	})

	t.Run("it should fail due to a small cep", func(t *testing.T) {
		cep, err := valueobjects.NewCep("01310")
		assert.Nil(t, cep)
		assert.NotNil(t, err)

		assert.Equal(t, "CEP inválido: 01310", err.Error())
	})

	t.Run("it should fail due to a big cep", func(t *testing.T) {
		cep, err := valueobjects.NewCep("013101001")
		assert.Nil(t, cep)
		assert.NotNil(t, err)

		assert.Equal(t, "CEP inválido: 013101001", err.Error())
	})

	t.Run("it should fail due to empty string", func(t *testing.T) {
		cep, err := valueobjects.NewCep("")
		assert.Nil(t, cep)
		assert.NotNil(t, err)

		assert.Equal(t, "CEP inválido: ", err.Error())
	})
}
