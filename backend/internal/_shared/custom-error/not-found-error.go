package customerror

import "fmt"

type NotFoundError struct {
	Key   string
	Value string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %s não encontrado", e.Key, e.Value)
}
