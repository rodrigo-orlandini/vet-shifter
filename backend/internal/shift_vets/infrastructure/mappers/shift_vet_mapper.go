package mappers

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
)

func ShiftVetFromPersistence(q queries.ShiftVet) (*entities.ShiftVet, error) {
	email, err := valueobjects.NewEmail(q.Email)
	if err != nil {
		return nil, err
	}
	phone, err := valueobjects.NewPhone(q.Phone)
	if err != nil {
		return nil, err
	}
	cpf, err := sharedvo.NewCpf(q.Cpf)
	if err != nil {
		return nil, err
	}
	v := &entities.ShiftVet{
		Id:             q.ID.String(),
		Email:          *email,
		Phone:          *phone,
		Password:       q.Password,
		FullName:       q.FullName,
		Cpf:            *cpf,
		CrmvNumber:     q.CrmvNumber,
		CrmvState:      q.CrmvState,
		Specialties:    q.Specialties,
		ApprovalStatus: q.ApprovalStatus,
		CreatedAt:      &q.CreatedAt,
	}
	if q.ConsentLgpdAt.Valid {
		v.ConsentLgpdAt = &q.ConsentLgpdAt.Time
	}
	if v.Specialties == nil {
		v.Specialties = []string{}
	}
	return v, nil
}

type ShiftVetFromHttpInput struct {
	Email         string
	Phone         string
	Password      string
	FullName      string
	Cpf           string
	CrmvNumber    string
	CrmvState     string
	Specialties   []string
	ConsentLgpdAt *time.Time
}

func ShiftVetFromHttp(input ShiftVetFromHttpInput) (*entities.ShiftVet, error) {
	email, err := valueobjects.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}
	phone, err := valueobjects.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	cpf, err := sharedvo.NewCpf(input.Cpf)
	if err != nil {
		return nil, err
	}
	if input.Specialties == nil {
		input.Specialties = []string{}
	}
	return entities.NewShiftVet(
		*email,
		*phone,
		input.Password,
		input.FullName,
		*cpf,
		input.CrmvNumber,
		input.CrmvState,
		input.Specialties,
		input.ConsentLgpdAt,
	)
}
