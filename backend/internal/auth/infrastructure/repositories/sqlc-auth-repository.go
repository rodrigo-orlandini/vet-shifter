package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/mappers"
)

type SqlcAuthRepository struct {
	queries *queries.Queries
}

func NewSqlcAuthRepository() *SqlcAuthRepository {
	conn := database.GetPostgresConnection()
	q := database.NewQueries(conn)
	return &SqlcAuthRepository{
		queries: q,
	}
}

func (r *SqlcAuthRepository) CreatePasswordResetToken(
	token string,
	email sharedvalueobjects.Email,
	userType sharedvalueobjects.UserType,
	expiresAt time.Time,
) (*entities.PasswordResetToken, error) {
	ctx := context.Background()
	id := uuid.Must(uuid.NewV7())

	created, err := r.queries.CreatePasswordResetToken(ctx, queries.CreatePasswordResetTokenParams{
		ID:        id,
		Token:     token,
		Email:     email.GetValue(),
		UserType:  queries.UserType(userType.GetValue()),
		ExpiresAt: expiresAt,
	})

	if err != nil {
		return nil, err
	}

	return mappers.PasswordResetTokenFromPersistence(created)
}

func (r *SqlcAuthRepository) GetPasswordResetToken(token string) (*entities.PasswordResetToken, error) {
	ctx := context.Background()
	t, err := r.queries.GetPasswordResetTokenByToken(ctx, token)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return mappers.PasswordResetTokenFromPersistence(t)
}

func (r *SqlcAuthRepository) MarkPasswordResetTokenUsed(id string) error {
	ctx := context.Background()

	return r.queries.MarkPasswordResetTokenUsed(ctx, uuid.MustParse(id))
}
