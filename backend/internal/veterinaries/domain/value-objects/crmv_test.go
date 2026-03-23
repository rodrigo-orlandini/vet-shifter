package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

func TestValueObjectCrmv(t *testing.T) {
	t.Run("it should be able to create valid crmv", func(t *testing.T) {
		crmv, err := valueobjects.NewCrmv("12345", "SP")
		assert.Nil(t, err)
		assert.NotNil(t, crmv)
		assert.Equal(t, "12345", crmv.GetNumber())
		assert.Equal(t, "SP", crmv.GetState())
	})

	t.Run("it should normalize state to uppercase", func(t *testing.T) {
		crmv, err := valueobjects.NewCrmv("99999", "rj")
		assert.Nil(t, err)
		assert.Equal(t, "RJ", crmv.GetState())
	})

	t.Run("it should fail for empty number", func(t *testing.T) {
		crmv, err := valueobjects.NewCrmv("", "SP")
		assert.Nil(t, crmv)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "CRMV")
	})

	t.Run("it should fail for invalid state", func(t *testing.T) {
		crmv, err := valueobjects.NewCrmv("12345", "XX")
		assert.Nil(t, crmv)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "CRMV")
	})

	t.Run("it should fail for empty state", func(t *testing.T) {
		crmv, err := valueobjects.NewCrmv("12345", "")
		assert.Nil(t, crmv)
		assert.NotNil(t, err)
	})
}
