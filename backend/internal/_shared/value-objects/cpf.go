package valueobjects

import (
	"fmt"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

type Cpf struct {
	value  string
	masked string
}

func NewCpf(value string) (*Cpf, error) {
	if len(value) != 11 {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Cpf",
			Value: value,
		}
	}
	masked := fmt.Sprintf("%s.%s.%s-%s", value[:3], value[3:6], value[6:9], value[9:])
	return &Cpf{value: value, masked: masked}, nil
}

func (c *Cpf) GetValue() string  { return c.value }
func (c *Cpf) GetMasked() string { return c.masked }
