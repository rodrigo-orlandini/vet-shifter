package repositories

import (
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
)

type ShiftVeterinaryRepository interface {
	Create(veterinary entities.ShiftVeterinary) (*entities.ShiftVeterinary, error)
	FindByCpf(cpf sharedvalueobjects.Cpf) (*entities.ShiftVeterinary, error)
	FindByEmail(email sharedvalueobjects.Email) (*entities.ShiftVeterinary, error)
	FindByPhone(phone sharedvalueobjects.Phone) (*entities.ShiftVeterinary, error)
	UpdatePassword(id string, hashedPassword string) error
}
