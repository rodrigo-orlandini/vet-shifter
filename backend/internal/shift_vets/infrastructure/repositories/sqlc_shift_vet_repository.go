package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	valueobjects "rodrigoorlandini/vet-shifter/internal/companies/domain/value-objects"
	sharedvo "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/mappers"
)

type SqlcShiftVetRepository struct {
	queries *queries.Queries
}

func NewSqlcShiftVetRepository() *SqlcShiftVetRepository {
	return &SqlcShiftVetRepository{queries: database.NewQueries(database.GetPostgresConnection())}
}

func (r *SqlcShiftVetRepository) Create(vet entities.ShiftVet) (*entities.ShiftVet, error) {
	ctx := context.Background()
	consentLgpdAt := sql.NullTime{}
	if vet.ConsentLgpdAt != nil {
		consentLgpdAt = sql.NullTime{Time: *vet.ConsentLgpdAt, Valid: true}
	}
	created, err := r.queries.CreateShiftVet(ctx, queries.CreateShiftVetParams{
		ID:             uuid.MustParse(vet.Id),
		Email:          vet.Email.GetValue(),
		Phone:          vet.Phone.GetValue(),
		Password:       vet.Password,
		FullName:       vet.FullName,
		Cpf:            vet.Cpf.GetValue(),
		CrmvNumber:     vet.CrmvNumber,
		CrmvState:      vet.CrmvState,
		Specialties:    vet.Specialties,
		ApprovalStatus: vet.ApprovalStatus,
		ConsentLgpdAt:  consentLgpdAt,
	})
	if err != nil {
		return nil, err
	}
	return mappers.ShiftVetFromPersistence(created)
}

func (r *SqlcShiftVetRepository) FindByEmail(email valueobjects.Email) (*entities.ShiftVet, error) {
	ctx := context.Background()
	q, err := r.queries.FindShiftVetByEmail(ctx, email.GetValue())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mappers.ShiftVetFromPersistence(q)
}

func (r *SqlcShiftVetRepository) FindByCpf(cpf sharedvo.Cpf) (*entities.ShiftVet, error) {
	ctx := context.Background()
	q, err := r.queries.FindShiftVetByCpf(ctx, cpf.GetValue())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mappers.ShiftVetFromPersistence(q)
}

func (r *SqlcShiftVetRepository) FindByID(id string) (*entities.ShiftVet, error) {
	ctx := context.Background()
	q, err := r.queries.FindShiftVetByID(ctx, uuid.MustParse(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return mappers.ShiftVetFromPersistence(q)
}
