package customerror

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return "Algo deu errado. Tente novamente."
}
