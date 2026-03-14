package mappers

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

func ShiftVeterinaryFromPersistence(t queries.ShiftVeterinary) (*entities.ShiftVeterinary, error) {
	email, err := sharedvalueobjects.NewEmail(t.Email)
	if err != nil {
		return nil, err
	}

	phone, err := sharedvalueobjects.NewPhone(t.Phone)
	if err != nil {
		return nil, err
	}

	cpf, err := sharedvalueobjects.NewCpf(t.Cpf)
	if err != nil {
		return nil, err
	}

	crmv, err := valueobjects.NewCrmv(t.CrmvNumber, t.CrmvState)
	if err != nil {
		return nil, err
	}

	specialtiesStrings := make([]string, 0, len(t.Specialties))
	for _, s := range t.Specialties {
		specialtiesStrings = append(specialtiesStrings, string(s))
	}

	specialties, err := valueobjects.NewSpecialties(specialtiesStrings)
	if err != nil {
		return nil, err
	}

	registrationStatus, err := valueobjects.NewRegistrationStatus(string(t.ApprovalStatus))
	if err != nil {
		return nil, err
	}

	var consentLgpdAt *time.Time
	if t.ConsentLgpdAt.Valid {
		consentLgpdAt = &t.ConsentLgpdAt.Time
	}

	return &entities.ShiftVeterinary{
		Id:                 t.ID.String(),
		Email:              *email,
		Phone:              *phone,
		Password:           t.Password,
		FullName:           t.FullName,
		Cpf:                *cpf,
		Crmv:               *crmv,
		Specialties:        *specialties,
		RegistrationStatus: *registrationStatus,
		ConsentLgpdAt:      consentLgpdAt,
		CreatedAt:          t.CreatedAt,
	}, nil
}
