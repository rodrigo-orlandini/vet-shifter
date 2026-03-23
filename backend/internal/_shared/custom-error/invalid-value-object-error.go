package customerror

import "fmt"

type InvalidValueObjectError struct {
	Key   string
	Value string
}

func (e *InvalidValueObjectError) Error() string {
	return fmt.Sprintf("%s inválido: %s", e.Key, e.Value)
}
