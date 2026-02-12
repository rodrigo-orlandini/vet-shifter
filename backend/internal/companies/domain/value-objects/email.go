package valueobjects

import (
	"regexp"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9]([a-zA-Z0-9.\-]*[a-zA-Z0-9])?\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(value) {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Email",
			Value: value,
		}
	}

	email := &Email{
		value: value,
	}

	return email, nil
}

func (c *Email) GetValue() string {
	return c.value
}
