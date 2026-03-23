package valueobjects

import (
	"fmt"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Phone struct {
	value  string
	masked string
}

func NewPhone(value string) (*Phone, error) {
	if len(value) != 11 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Telefone",
			Value: value,
		}
	}

	masked := fmt.Sprintf("(%s) %s-%s", value[:2], value[2:7], value[7:])

	phone := &Phone{
		value:  value,
		masked: masked,
	}

	return phone, nil
}

func (c *Phone) GetValue() string {
	return c.value
}

func (c *Phone) GetMasked() string {
	return c.masked
}
