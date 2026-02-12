package entities

import (
	"time"

	"github.com/google/uuid"
)

const (
	ShiftStatusOpen       = "open"
	ShiftStatusConfirmed  = "confirmed"
	ShiftStatusCancelled  = "cancelled"
	ShiftStatusCompleted  = "completed"
)

const (
	ShiftTypeEmergency   = "emergency"
	ShiftTypeConsultation = "consultation"
	ShiftTypeSurgery     = "surgery"
)

type Shift struct {
	Id                string
	CompanyId         string
	StartsAt          time.Time
	EndsAt            time.Time
	Type              string
	OfferedValueCents int64
	Requirements      string
	Description       string
	Location          string
	Status            string
	CreatedAt         time.Time
}

func NewShift(
	companyId string,
	startsAt, endsAt time.Time,
	shiftType string,
	offeredValueCents int64,
	requirements, description, location string,
) (*Shift, error) {
	id, _ := uuid.NewV7()
	now := time.Now()
	return &Shift{
		Id:                id.String(),
		CompanyId:         companyId,
		StartsAt:          startsAt,
		EndsAt:            endsAt,
		Type:              shiftType,
		OfferedValueCents: offeredValueCents,
		Requirements:      requirements,
		Description:       description,
		Location:          location,
		Status:            ShiftStatusOpen,
		CreatedAt:         now,
	}, nil
}
