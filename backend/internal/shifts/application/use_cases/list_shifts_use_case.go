package usecases

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/shifts/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
)

type ListShiftsUseCase struct {
	repo repositories.ShiftRepository
}

type ListShiftsInput struct {
	Status    string
	Type      string
	CompanyID string
	FromDate  string
	ToDate    string
	Limit     int32
	Offset    int32
}

type ListShiftsOutput struct {
	Shifts []*entities.Shift
}

func NewListShiftsUseCase(repo repositories.ShiftRepository) *ListShiftsUseCase {
	return &ListShiftsUseCase{repo: repo}
}

func (u *ListShiftsUseCase) Execute(input *ListShiftsInput) (*ListShiftsOutput, error) {
	filters := repositories.ListShiftsFilters{
		Status:    input.Status,
		Type:      input.Type,
		CompanyID: input.CompanyID,
		Limit:     input.Limit,
		Offset:    input.Offset,
	}
	if input.FromDate != "" {
		if t, err := parseTime(input.FromDate); err == nil {
			filters.FromDate = t
		}
	}
	if input.ToDate != "" {
		if t, err := parseTime(input.ToDate); err == nil {
			filters.ToDate = t
		}
	}
	list, err := u.repo.List(filters)
	if err != nil {
		return nil, err
	}
	return &ListShiftsOutput{Shifts: list}, nil
}

func parseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}
