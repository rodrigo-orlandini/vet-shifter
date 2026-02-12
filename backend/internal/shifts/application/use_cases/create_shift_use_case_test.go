package usecases_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	usecases "rodrigoorlandini/vet-shifter/internal/shifts/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func TestCreateShiftUseCase(t *testing.T) {
	t.Run("creates shift successfully", func(t *testing.T) {
		repo := repositories.NewStubShiftRepository()
		uc := usecases.NewCreateShiftUseCase(repo)
		starts := time.Now().Add(24 * time.Hour)
		ends := starts.Add(8 * time.Hour)
		shift, _ := entities.NewShift(
			"550e8400-e29b-41d4-a716-446655440000",
			starts, ends,
			entities.ShiftTypeConsultation,
			15000,
			"", "", "",
		)
		out, err := uc.Execute(&usecases.CreateShiftInput{Shift: *shift})
		assert.Nil(t, err)
		assert.NotNil(t, out.Shift)
		assert.NotEmpty(t, out.Shift.Id)
		assert.Equal(t, entities.ShiftStatusOpen, out.Shift.Status)
	})
}
