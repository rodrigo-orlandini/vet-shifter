package valueobjects

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

const (
	UserTypeCompanyOwner    = "company_owner"
	UserTypeShiftVeterinary = "shift_veterinary"
)

type UserType struct {
	value string
}

func NewUserType(value string) (*UserType, error) {
	if value != UserTypeCompanyOwner && value != UserTypeShiftVeterinary {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Tipo de Usuário",
			Value: value,
		}
	}

	return &UserType{value: value}, nil
}

func (u *UserType) GetValue() string {
	return u.value
}

func (u *UserType) Equals(other *UserType) bool {
	if other == nil {
		return false
	}

	return u.value == other.value
}

func CompanyOwner() *UserType {
	return &UserType{value: UserTypeCompanyOwner}
}

func ShiftVeterinary() *UserType {
	return &UserType{value: UserTypeShiftVeterinary}
}
