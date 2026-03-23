package valueobjects

import (
	"fmt"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Cep struct {
	value  string
	masked string
}

func NewCep(value string) (*Cep, error) {
	if len(value) != 8 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "CEP",
			Value: value,
		}
	}

	masked := fmt.Sprintf("%s-%s", value[:5], value[5:])

	return &Cep{
		value:  value,
		masked: masked,
	}, nil
}

func (c *Cep) GetValue() string {
	return c.value
}

func (c *Cep) GetMasked() string {
	return c.masked
}
