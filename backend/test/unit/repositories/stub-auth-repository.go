package repositories

import (
	"time"

	"github.com/google/uuid"

	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	"rodrigoorlandini/vet-shifter/internal/auth/domain/entities"
)

type StubAuthRepository struct {
	users  map[string]*entities.GenericUser
	tokens map[string]*entities.PasswordResetToken
}

func NewStubAuthRepository() *StubAuthRepository {
	return &StubAuthRepository{
		users:  make(map[string]*entities.GenericUser),
		tokens: make(map[string]*entities.PasswordResetToken),
	}
}

func (r *StubAuthRepository) AddUser(
	id string,
	email sharedvalueobjects.Email,
	userType sharedvalueobjects.UserType,
	passwordHash string,
) {
	u, _ := entities.NewGenericUser(id, email, passwordHash, userType)
	r.users[id] = u
}

func (r *StubAuthRepository) CreatePasswordResetToken(
	token string,
	email sharedvalueobjects.Email,
	userType sharedvalueobjects.UserType,
	expiresAt time.Time,
) (*entities.PasswordResetToken, error) {
	id := uuid.Must(uuid.NewV7()).String()
	rec, err := entities.NewPasswordResetToken(id, token, email, userType, expiresAt, nil)
	if err != nil {
		return nil, err
	}

	r.tokens[token] = rec
	return rec, nil
}

func (r *StubAuthRepository) GetPasswordResetToken(token string) (*entities.PasswordResetToken, error) {
	if t, ok := r.tokens[token]; ok {
		return t, nil
	}

	return nil, nil
}

func (r *StubAuthRepository) MarkPasswordResetTokenUsed(id string) error {
	for token, t := range r.tokens {
		if t.Id == id {
			now := time.Now()
			updated, _ := entities.NewPasswordResetToken(t.Id, t.Token, t.Email, t.UserType, t.ExpiresAt, &now)
			r.tokens[token] = updated
			break
		}
	}

	return nil
}
