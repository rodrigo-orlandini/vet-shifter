package customerror

import "fmt"

type AlreadyExistsError struct {
	Field string
	Value string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s já cadastrado: %s", e.Field, e.Value)
}
