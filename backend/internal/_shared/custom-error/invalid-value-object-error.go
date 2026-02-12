package customerror

import "fmt"

type InvalidValueObjectError struct {
	Key   string
	Value string
}

func (e *InvalidValueObjectError) Error() string {
	return fmt.Sprintf("Invalid value object '%s' creation with value: %s", e.Key, e.Value)
}
