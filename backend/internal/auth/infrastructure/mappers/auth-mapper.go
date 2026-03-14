package mappers

import (
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
)

func PasswordResetTokenFromPersistence(t queries.PasswordResetToken) (*entities.PasswordResetToken, error) {
	emailVO, err := sharedvalueobjects.NewEmail(t.Email)
	if err != nil {
		return nil, err
	}

	userTypeVO, err := sharedvalueobjects.NewUserType(string(t.UserType))
	if err != nil {
		return nil, err
	}

	var usedAt *time.Time
	if t.UsedAt.Valid {
		usedAt = &t.UsedAt.Time
	}

	return entities.NewPasswordResetToken(t.ID.String(), t.Token, *emailVO, *userTypeVO, t.ExpiresAt, usedAt)
}
