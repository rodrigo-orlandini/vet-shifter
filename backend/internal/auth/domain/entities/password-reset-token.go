package entities

import (
	"time"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type PasswordResetToken struct {
	Id        string
	Token     string
	Email     sharedvalueobjects.Email
	UserType  sharedvalueobjects.UserType
	ExpiresAt time.Time
	UsedAt    *time.Time
}

func NewPasswordResetToken(
	id string,
	token string,
	email sharedvalueobjects.Email,
	userType sharedvalueobjects.UserType,
	expiresAt time.Time,
	usedAt *time.Time,
) (*PasswordResetToken, error) {
	return &PasswordResetToken{
		Id:        id,
		Token:     token,
		Email:     email,
		UserType:  userType,
		ExpiresAt: expiresAt,
		UsedAt:    usedAt,
	}, nil
}
