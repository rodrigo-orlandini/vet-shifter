package customerror

import "fmt"

type AlreadyExistsError struct {
	Entity string
	Field  string
	Value  string
}

func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("The value '%s' already exists for field '%s' of entity '%s'", e.Value, e.Field, e.Entity)
}
