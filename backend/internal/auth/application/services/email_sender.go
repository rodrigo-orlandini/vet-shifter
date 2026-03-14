package services

import valueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"

type EmailSender interface {
	SendPasswordResetEmail(to valueobjects.Email, resetLink string) error
}
