package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	shiftrepos "rodrigoorlandini/vet-shifter/internal/shifts/application/repositories"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/shifts/infrastructure/mappers"
)

type SqlcShiftRepository struct {
	queries *queries.Queries
}

func NewSqlcShiftRepository() *SqlcShiftRepository {
	return &SqlcShiftRepository{queries: database.NewQueries(database.GetPostgresConnection())}
}

func (r *SqlcShiftRepository) Create(shift entities.Shift) (*entities.Shift, error) {
	ctx := context.Background()
	req := sql.NullString{}
	if shift.Requirements != "" {
		req = sql.NullString{String: shift.Requirements, Valid: true}
	}
	desc := sql.NullString{}
	if shift.Description != "" {
		desc = sql.NullString{String: shift.Description, Valid: true}
	}
	loc := sql.NullString{}
	if shift.Location != "" {
		loc = sql.NullString{String: shift.Location, Valid: true}
	}
	created, err := r.queries.CreateShift(ctx, queries.CreateShiftParams{
		ID:                uuid.MustParse(shift.Id),
		CompanyID:         uuid.MustParse(shift.CompanyId),
		StartsAt:          shift.StartsAt,
		EndsAt:            shift.EndsAt,
		Type:              shift.Type,
		OfferedValueCents: shift.OfferedValueCents,
		Requirements:      req,
		Description:       desc,
		Location:          loc,
		Status:            shift.Status,
	})
	if err != nil {
		return nil, err
	}
	return mappers.ShiftFromPersistence(created), nil
}

func (r *SqlcShiftRepository) GetByID(id string) (*entities.Shift, error) {
	ctx := context.Background()
	q, err := r.queries.GetShiftByID(ctx, uuid.MustParse(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mappers.ShiftFromPersistence(q), nil
}

func (r *SqlcShiftRepository) List(filters shiftrepos.ListShiftsFilters) ([]*entities.Shift, error) {
	ctx := context.Background()
	if filters.Limit <= 0 {
		filters.Limit = 50
	}
	fromDate := filters.FromDate
	if fromDate.IsZero() {
		fromDate = time.Time{}
	}
	toDate := filters.ToDate
	if toDate.IsZero() {
		toDate = time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
	}
	companyID := uuid.Nil
	if filters.CompanyID != "" {
		companyID = uuid.MustParse(filters.CompanyID)
	}
	list, err := r.queries.ListShifts(ctx, queries.ListShiftsParams{
		Column1:    filters.Status,
		Column2:    filters.Type,
		Column3:    companyID,
		StartsAt:   fromDate,
		StartsAt_2: toDate,
		Limit:      filters.Limit,
		Offset:     filters.Offset,
	})
	if err != nil {
		return nil, err
	}
	out := make([]*entities.Shift, len(list))
	for i := range list {
		out[i] = mappers.ShiftFromPersistence(list[i])
	}
	return out, nil
}
