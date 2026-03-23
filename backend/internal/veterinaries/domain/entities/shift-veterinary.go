package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	valueobjects "rodrigoorlandini/vet-shifter/internal/veterinaries/domain/value-objects"
)

const MinFullNameLength = 2

type ShiftVeterinary struct {
	Id                 string
	Email              sharedvalueobjects.Email
	Phone              sharedvalueobjects.Phone
	Password           string
	FullName           string
	Cpf                sharedvalueobjects.Cpf
	Crmv               valueobjects.Crmv
	Specialties        valueobjects.Specialties
	RegistrationStatus valueobjects.RegistrationStatus
	ConsentLgpdAt      *time.Time
	CreatedAt          time.Time
}

func NewShiftVeterinary(
	email sharedvalueobjects.Email,
	phone sharedvalueobjects.Phone,
	password string,
	fullName string,
	cpf sharedvalueobjects.Cpf,
	crmv valueobjects.Crmv,
	specialties valueobjects.Specialties,
	registrationStatus valueobjects.RegistrationStatus,
	consentLgpdAt *time.Time,
) (*ShiftVeterinary, error) {
	if consentLgpdAt == nil {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Consentimento LGPD",
			Value: "",
		}
	}

	trimmedName := strings.TrimSpace(fullName)
	if len(trimmedName) < MinFullNameLength {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Nome completo",
			Value: fullName,
		}
	}

	if len(password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Senha",
			Value: "",
		}
	}

	id, _ := uuid.NewV7()

	return &ShiftVeterinary{
		Id:                 id.String(),
		Email:              email,
		Phone:              phone,
		Password:           password,
		FullName:           trimmedName,
		Cpf:                cpf,
		Crmv:               crmv,
		Specialties:        specialties,
		RegistrationStatus: registrationStatus,
		ConsentLgpdAt:      consentLgpdAt,
		CreatedAt:          time.Now(),
	}, nil
}
