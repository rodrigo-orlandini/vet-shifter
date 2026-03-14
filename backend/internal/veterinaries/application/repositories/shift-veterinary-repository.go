package repositories

import (
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
)

type ShiftVeterinaryRepository interface {
	FindByEmail(email sharedvalueobjects.Email) (*entities.ShiftVeterinary, error)
	UpdatePassword(id string, hashedPassword string) error
}
