package valueobjects_test

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueObjectPhone(t *testing.T) {
	t.Run("it should be able to create new phone", func(t *testing.T) {
		phone, err := valueobjects.NewPhone("00000000000")
		assert.Nil(t, err)

		assert.Equal(t, phone.GetValue(), "00000000000")
		assert.Equal(t, phone.GetMasked(), "(00) 00000-0000")
	})

	t.Run("it should fail due to a small phone", func(t *testing.T) {
		phone, err := valueobjects.NewPhone("0")
		assert.Nil(t, phone)
		assert.NotNil(t, err)

		assert.Equal(t, err.Error(), "Invalid value object 'Phone' creation with value: 0")
	})

	t.Run("it should fail due to a big phone", func(t *testing.T) {
		phone, err := valueobjects.NewPhone("000000000000000")
		assert.Nil(t, phone)
		assert.NotNil(t, err)

		assert.Equal(t, err.Error(), "Invalid value object 'Phone' creation with value: 000000000000000")
	})
}
