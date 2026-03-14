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

func (r *SqlcShiftVeterinaryRepository) UpdatePassword(userID string, hashedPassword string) error {
	ctx := context.Background()
	id := uuid.MustParse(userID)

	return r.queries.UpdateShiftVeterinaryPassword(ctx, queries.UpdateShiftVeterinaryPasswordParams{
		ID:       id,
		Password: hashedPassword,
	})
}
