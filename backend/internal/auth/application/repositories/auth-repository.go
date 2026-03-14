package repositories

import (
	"time"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
)

type AuthRepository interface {
	CreatePasswordResetToken(
		token string,
		email sharedvalueobjects.Email,
		userType sharedvalueobjects.UserType,
		expiresAt time.Time,
	) (*entities.PasswordResetToken, error)
	GetPasswordResetToken(token string) (*entities.PasswordResetToken, error)
	MarkPasswordResetTokenUsed(id string) error
}
