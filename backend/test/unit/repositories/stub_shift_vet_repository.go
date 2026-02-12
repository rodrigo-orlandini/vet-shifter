package repositories

import (
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
)

type StubShiftVetRepository struct {
	vets map[string]*entities.ShiftVet
	byEmail map[string]*entities.ShiftVet
	byCpf  map[string]*entities.ShiftVet
}

func NewStubShiftVetRepository() *StubShiftVetRepository {
	return &StubShiftVetRepository{
		vets:    make(map[string]*entities.ShiftVet),
		byEmail: make(map[string]*entities.ShiftVet),
		byCpf:   make(map[string]*entities.ShiftVet),
	}
}

func (r *StubShiftVetRepository) Create(vet entities.ShiftVet) (*entities.ShiftVet, error) {
	r.vets[vet.Id] = &vet
	r.byEmail[vet.Email.GetValue()] = &vet
	r.byCpf[vet.Cpf.GetValue()] = &vet
	return &vet, nil
}

func (r *StubShiftVetRepository) FindByEmail(email valueobjects.Email) (*entities.ShiftVet, error) {
	return r.byEmail[email.GetValue()], nil
}

func (r *StubShiftVetRepository) FindByCpf(cpf sharedvo.Cpf) (*entities.ShiftVet, error) {
	return r.byCpf[cpf.GetValue()], nil
}

func (r *StubShiftVetRepository) FindByID(id string) (*entities.ShiftVet, error) {
	return r.vets[id], nil
}
