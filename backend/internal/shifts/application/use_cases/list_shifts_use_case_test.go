package usecases_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	usecases "rodrigoorlandini/vet-shifter/internal/shifts/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
	"rodrigoorlandini/vet-shifter/test/unit/repositories"
)

func TestListShiftsUseCase(t *testing.T) {
	t.Run("returns shifts matching filters", func(t *testing.T) {
		repo := repositories.NewStubShiftRepository()
		uc := usecases.NewListShiftsUseCase(repo)
		starts := time.Now().Add(24 * time.Hour)
		ends := starts.Add(8 * time.Hour)
		shift, _ := entities.NewShift("company-1", starts, ends, entities.ShiftTypeEmergency, 20000, "", "", "")
		repo.Create(*shift)
		out, err := uc.Execute(&usecases.ListShiftsInput{Limit: 10, Offset: 0})
		assert.Nil(t, err)
		assert.Len(t, out.Shifts, 1)
		assert.Equal(t, entities.ShiftTypeEmergency, out.Shifts[0].Type)
	})
}
