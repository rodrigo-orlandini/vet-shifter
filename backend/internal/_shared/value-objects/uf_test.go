package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

func TestValueObjectUF(t *testing.T) {
	t.Run("it should be able to create valid UFs", func(t *testing.T) {
		cases := []string{
			"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES",
			"GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR",
			"PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC",
			"SP", "SE", "TO",
		}

		for _, uf := range cases {
			result, err := valueobjects.NewUF(uf)
			assert.Nil(t, err, "expected no error for UF: %s", uf)
			assert.Equal(t, uf, result.GetValue())
		}
	})

	t.Run("it should normalize lowercase to uppercase", func(t *testing.T) {
		uf, err := valueobjects.NewUF("sp")
		assert.Nil(t, err)
		assert.Equal(t, "SP", uf.GetValue())
	})

	t.Run("it should trim whitespace", func(t *testing.T) {
		uf, err := valueobjects.NewUF(" RJ ")
		assert.Nil(t, err)
		assert.Equal(t, "RJ", uf.GetValue())
	})

	t.Run("it should fail for invalid UF code", func(t *testing.T) {
		uf, err := valueobjects.NewUF("XX")
		assert.Nil(t, uf)
		assert.NotNil(t, err)
		assert.Equal(t, "UF inválido: XX", err.Error())
	})

	t.Run("it should fail for single character", func(t *testing.T) {
		uf, err := valueobjects.NewUF("S")
		assert.Nil(t, uf)
		assert.NotNil(t, err)
	})

	t.Run("it should fail for three characters", func(t *testing.T) {
		uf, err := valueobjects.NewUF("SPP")
		assert.Nil(t, uf)
		assert.NotNil(t, err)
	})

	t.Run("it should fail for empty string", func(t *testing.T) {
		uf, err := valueobjects.NewUF("")
		assert.Nil(t, uf)
		assert.NotNil(t, err)
		assert.Equal(t, "UF inválido: ", err.Error())
	})
}
