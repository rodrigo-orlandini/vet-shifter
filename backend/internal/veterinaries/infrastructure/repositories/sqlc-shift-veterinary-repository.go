package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/mappers"
)

type SqlcShiftVeterinaryRepository struct {
	queries *queries.Queries
}

func NewSqlcShiftVeterinaryRepository() *SqlcShiftVeterinaryRepository {
	conn := database.GetPostgresConnection()
	q := database.NewQueries(conn)

	return &SqlcShiftVeterinaryRepository{queries: q}
}

func (r *SqlcShiftVeterinaryRepository) Create(veterinary entities.ShiftVeterinary) (*entities.ShiftVeterinary, error) {
	ctx := context.Background()

	consentLgpdAt := sql.NullTime{}
	if veterinary.ConsentLgpdAt != nil {
		consentLgpdAt = sql.NullTime{Time: *veterinary.ConsentLgpdAt, Valid: true}
	}

	specialties := make([]queries.VeterinarySpecialty, 0, veterinary.Specialties.Len())
	for _, s := range veterinary.Specialties.GetValue() {
		specialties = append(specialties, queries.VeterinarySpecialty(s))
	}

	created, err := r.queries.CreateShiftVeterinary(ctx, queries.CreateShiftVeterinaryParams{
		ID:             uuid.MustParse(veterinary.Id),
		Email:          veterinary.Email.GetValue(),
		Phone:          veterinary.Phone.GetValue(),
		Password:       veterinary.Password,
		FullName:       veterinary.FullName,
		Cpf:            veterinary.Cpf.GetValue(),
		CrmvNumber:     veterinary.Crmv.GetNumber(),
		CrmvState:      veterinary.Crmv.GetState(),
		Specialties:    specialties,
		ApprovalStatus: queries.AccountStatus(veterinary.RegistrationStatus.String()),
		ConsentLgpdAt:  consentLgpdAt,
	})
	if err != nil {
		return nil, err
	}

	return mappers.ShiftVeterinaryFromPersistence(created)
}

func (r *SqlcShiftVeterinaryRepository) FindByCpf(
	cpf sharedvalueobjects.Cpf,
) (*entities.ShiftVeterinary, error) {
	ctx := context.Background()
	veterinary, err := r.queries.FindShiftVeterinaryByCpf(ctx, cpf.GetValue())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return mappers.ShiftVeterinaryFromPersistence(veterinary)
}

func (r *SqlcShiftVeterinaryRepository) FindByEmail(
	email sharedvalueobjects.Email,
) (*entities.ShiftVeterinary, error) {
	ctx := context.Background()
	veterinary, err := r.queries.FindShiftVeterinaryByEmail(ctx, email.GetValue())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return mappers.ShiftVeterinaryFromPersistence(veterinary)
}

func (r *SqlcShiftVeterinaryRepository) FindByPhone(
	phone sharedvalueobjects.Phone,
) (*entities.ShiftVeterinary, error) {
	ctx := context.Background()
	veterinary, err := r.queries.FindShiftVeterinaryByPhone(ctx, phone.GetValue())

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return mappers.ShiftVeterinaryFromPersistence(veterinary)
}

func (r *SqlcShiftVeterinaryRepository) UpdatePassword(userID string, hashedPassword string) error {
	ctx := context.Background()
	id := uuid.MustParse(userID)

	return r.queries.UpdateShiftVeterinaryPassword(ctx, queries.UpdateShiftVeterinaryPasswordParams{
		ID:       id,
		Password: hashedPassword,
	})
}
