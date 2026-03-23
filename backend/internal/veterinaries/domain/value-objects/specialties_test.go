package valueobjects_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

func TestValueObjectSpecialties(t *testing.T) {
	t.Run("it should create with a single specialty", func(t *testing.T) {
		sp, err := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyGeneralPractice})
		assert.Nil(t, err)
		assert.NotNil(t, sp)
		assert.Equal(t, 1, sp.Len())
		assert.Equal(t, []string{valueobjects.SpecialtyGeneralPractice}, sp.GetValue())
	})

	t.Run("it should create with multiple specialties", func(t *testing.T) {
		input := []string{
			valueobjects.SpecialtyCardiology,
			valueobjects.SpecialtyNeurology,
			valueobjects.SpecialtyICU,
		}

		sp, err := valueobjects.NewSpecialties(input)
		assert.Nil(t, err)
		assert.Equal(t, 3, sp.Len())
		assert.Equal(t, input, sp.GetValue())
	})

	t.Run("it should fail for empty list", func(t *testing.T) {
		sp, err := valueobjects.NewSpecialties([]string{})
		assert.Nil(t, sp)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Especialidades")
	})

	t.Run("it should fail for nil list", func(t *testing.T) {
		sp, err := valueobjects.NewSpecialties(nil)
		assert.Nil(t, sp)
		assert.NotNil(t, err)
	})

	t.Run("it should fail for invalid specialty", func(t *testing.T) {
		sp, err := valueobjects.NewSpecialties([]string{"invalid_specialty"})
		assert.Nil(t, sp)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Especialidades")
	})

	t.Run("it should fail if any specialty in the list is invalid", func(t *testing.T) {
		sp, err := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyFelines, "made_up"})
		assert.Nil(t, sp)
		assert.NotNil(t, err)
	})

	t.Run("GetValue should return a copy", func(t *testing.T) {
		sp, _ := valueobjects.NewSpecialties([]string{valueobjects.SpecialtyDermatology})
		values := sp.GetValue()
		values[0] = "tampered"
		assert.Equal(t, valueobjects.SpecialtyDermatology, sp.GetValue()[0])
	})

	t.Run("AllAvailableSpecialties should return all specialties", func(t *testing.T) {
		all := valueobjects.AllAvailableSpecialties()
		assert.Equal(t, 20, len(all))
	})
}
