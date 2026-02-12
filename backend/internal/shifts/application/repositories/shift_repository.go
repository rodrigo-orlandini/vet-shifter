package repositories

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
)

type ShiftRepository interface {
	Create(shift entities.Shift) (*entities.Shift, error)
	GetByID(id string) (*entities.Shift, error)
	List(filters ListShiftsFilters) ([]*entities.Shift, error)
}

type ListShiftsFilters struct {
	Status    string
	Type      string
	CompanyID string
	FromDate  time.Time
	ToDate    time.Time
	Limit     int32
	Offset    int32
}
