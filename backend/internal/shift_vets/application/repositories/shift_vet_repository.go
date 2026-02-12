package repositories

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
)

type ShiftVetRepository interface {
	Create(vet entities.ShiftVet) (*entities.ShiftVet, error)
	FindByEmail(email valueobjects.Email) (*entities.ShiftVet, error)
	FindByCpf(cpf sharedvo.Cpf) (*entities.ShiftVet, error)
	FindByID(id string) (*entities.ShiftVet, error)
}
