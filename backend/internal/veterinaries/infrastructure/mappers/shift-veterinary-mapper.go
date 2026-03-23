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

type ShiftVeterinaryFromHttpInput struct {
	Email        string
	Phone        string
	Password     string
	FullName     string
	Cpf          string
	CrmvNumber   string
	CrmvState    string
	Specialties  []string
	ConsentLgpd  bool
	ConsentLgpdAt *time.Time
}

func ShiftVeterinaryFromHttp(input ShiftVeterinaryFromHttpInput) (*entities.ShiftVeterinary, error) {
	email, err := sharedvalueobjects.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	phone, err := sharedvalueobjects.NewPhone(input.Phone)
	if err != nil {
		return nil, err
	}

	cpf, err := sharedvalueobjects.NewCpf(input.Cpf)
	if err != nil {
		return nil, err
	}

	crmv, err := valueobjects.NewCrmv(input.CrmvNumber, input.CrmvState)
	if err != nil {
		return nil, err
	}

	specialties, err := valueobjects.NewSpecialties(input.Specialties)
	if err != nil {
		return nil, err
	}

	var consentAt *time.Time
	if input.ConsentLgpdAt != nil {
		consentAt = input.ConsentLgpdAt
	} else if input.ConsentLgpd {
		t := time.Now()
		consentAt = &t
	}

	return entities.NewShiftVeterinary(
		*email,
		*phone,
		input.Password,
		input.FullName,
		*cpf,
		*crmv,
		*specialties,
		*valueobjects.PendingDocumentApproval(),
		consentAt,
	)
}
