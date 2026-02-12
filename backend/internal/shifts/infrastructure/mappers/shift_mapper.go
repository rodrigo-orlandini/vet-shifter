package mappers

import (
	"database/sql"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
)

func ShiftFromPersistence(q queries.Shift) *entities.Shift {
	s := &entities.Shift{
		Id:                q.ID.String(),
		CompanyId:         q.CompanyID.String(),
		StartsAt:          q.StartsAt,
		EndsAt:            q.EndsAt,
		Type:              q.Type,
		OfferedValueCents: q.OfferedValueCents,
		Requirements:      nullStringToStr(q.Requirements),
		Description:       nullStringToStr(q.Description),
		Location:          nullStringToStr(q.Location),
		Status:            q.Status,
		CreatedAt:         q.CreatedAt,
	}
	return s
}

func nullStringToStr(n sql.NullString) string {
	if n.Valid {
		return n.String
	}
	return ""
}
