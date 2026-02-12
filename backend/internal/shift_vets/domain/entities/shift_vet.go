package entities

import (
	"time"

	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"

	"github.com/google/uuid"
)

const (
	ShiftVetApprovalPending  = "pending"
	ShiftVetApprovalApproved = "approved"
	ShiftVetApprovalRejected = "rejected"
)

type ShiftVet struct {
	Id             string
	Email          valueobjects.Email
	Phone          valueobjects.Phone
	Password       string
	FullName       string
	Cpf            sharedvo.Cpf
	CrmvNumber     string
	CrmvState      string
	Specialties    []string
	ApprovalStatus string
	ConsentLgpdAt  *time.Time
	CreatedAt      *time.Time
}

func NewShiftVet(
	email valueobjects.Email,
	phone valueobjects.Phone,
	password string,
	fullName string,
	cpf sharedvo.Cpf,
	crmvNumber, crmvState string,
	specialties []string,
	consentLgpdAt *time.Time,
) (*ShiftVet, error) {
	now := time.Now()
	id, _ := uuid.NewV7()
	if specialties == nil {
		specialties = []string{}
	}
	return &ShiftVet{
		Id:             id.String(),
		Email:          email,
		Phone:          phone,
		Password:       password,
		FullName:       fullName,
		Cpf:            cpf,
		CrmvNumber:     crmvNumber,
		CrmvState:      crmvState,
		Specialties:    specialties,
		ApprovalStatus: ShiftVetApprovalPending,
		ConsentLgpdAt:  consentLgpdAt,
		CreatedAt:      &now,
	}, nil
}
