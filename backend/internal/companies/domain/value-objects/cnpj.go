package valueobjects

import (
	"fmt"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Cnpj struct {
	value  string
	masked string
}

func NewCnpj(value string) (*Cnpj, error) {
	if len(value) != 14 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Cnpj",
			Value: value,
		}
	}

	masked := fmt.Sprintf("%s.%s.%s/%s-%s", value[:2], value[2:5], value[5:8], value[8:12], value[12:])

	cnpj := &Cnpj{
		value:  value,
		masked: masked,
	}

	return cnpj, nil
}

func (c *Cnpj) GetValue() string {
	return c.value
}

func (c *Cnpj) GetMasked() string {
	return c.masked
}
