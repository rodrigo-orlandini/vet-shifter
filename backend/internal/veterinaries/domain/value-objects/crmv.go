package valueobjects

import (
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
)

type Crmv struct {
	number string
	state  sharedvalueobjects.UF
}

func NewCrmv(number string, state string) (*Crmv, error) {
	if number == "" {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "Número do CRMV",
			Value: number,
		}
	}

	uf, err := sharedvalueobjects.NewUF(state)
	if err != nil {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "UF do CRMV",
			Value: state,
		}
	}

	return &Crmv{number: number, state: *uf}, nil
}

func (c *Crmv) GetNumber() string {
	return c.number
}

func (c *Crmv) GetState() string {
	return c.state.GetValue()
}
