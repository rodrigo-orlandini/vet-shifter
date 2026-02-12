package repositories

import (
	shiftrepos "rodrigoorlandini/vet-shifter/internal/shifts/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
)

type StubShiftRepository struct {
	shifts map[string]*entities.Shift
}

func NewStubShiftRepository() *StubShiftRepository {
	return &StubShiftRepository{shifts: make(map[string]*entities.Shift)}
}

func (r *StubShiftRepository) Create(shift entities.Shift) (*entities.Shift, error) {
	r.shifts[shift.Id] = &shift
	return &shift, nil
}

func (r *StubShiftRepository) GetByID(id string) (*entities.Shift, error) {
	return r.shifts[id], nil
}

func (r *StubShiftRepository) List(filters shiftrepos.ListShiftsFilters) ([]*entities.Shift, error) {
	var out []*entities.Shift
	for _, s := range r.shifts {
		if filters.Status != "" && s.Status != filters.Status {
			continue
		}
		if filters.Type != "" && s.Type != filters.Type {
			continue
		}
		if filters.CompanyID != "" && s.CompanyId != filters.CompanyID {
			continue
		}
		out = append(out, s)
	}
	return out, nil
}
