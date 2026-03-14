package repositories

import (
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type StubEmailSender struct {
	SentEmails []SentPasswordResetEmail
}

type SentPasswordResetEmail struct {
	To       sharedvalueobjects.Email
	ResetLink string
}

func NewStubEmailSender() *StubEmailSender {
	return &StubEmailSender{
		SentEmails: nil,
	}
}

func (s *StubEmailSender) SendPasswordResetEmail(to sharedvalueobjects.Email, resetLink string) error {
	s.SentEmails = append(s.SentEmails, SentPasswordResetEmail{To: to, ResetLink: resetLink})
	return nil
}
