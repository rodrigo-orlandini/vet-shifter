package valueobjects

import (
	"strings"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
)

var validUFs = map[string]bool{
	"AC": true, "AL": true, "AP": true, "AM": true,
	"BA": true, "CE": true, "DF": true, "ES": true,
	"GO": true, "MA": true, "MT": true, "MS": true,
	"MG": true, "PA": true, "PB": true, "PR": true,
	"PE": true, "PI": true, "RJ": true, "RN": true,
	"RS": true, "RO": true, "RR": true, "SC": true,
	"SP": true, "SE": true, "TO": true,
}

type UF struct {
	value string
}

func NewUF(value string) (*UF, error) {
	upper := strings.ToUpper(strings.TrimSpace(value))

	if !validUFs[upper] {
		return nil, &customerror.InvalidValueObjectError{
			Key:   "UF",
			Value: value,
		}
	}

	return &UF{value: upper}, nil
}

func (u *UF) GetValue() string {
	return u.value
}
