package entities

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type GenericUser struct {
	Id       string
	Email    sharedvalueobjects.Email
	Password string
	Type     sharedvalueobjects.UserType
}

func NewGenericUser(
	id string,
	email sharedvalueobjects.Email,
	password string,
	userType sharedvalueobjects.UserType,
) (*GenericUser, error) {
	if len(password) < utils.MinPasswordLength {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Password",
			Value: "",
		}
	}

	return &GenericUser{
		Id:       id,
		Email:    email,
		Password: password,
		Type:     userType,
	}, nil
}
