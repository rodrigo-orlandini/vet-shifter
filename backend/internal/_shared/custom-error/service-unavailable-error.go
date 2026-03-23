package customerror

type ServiceUnavailableError struct {
	Service string
	Err     error
}

func (e *ServiceUnavailableError) Error() string {
	return "Serviço temporariamente indisponível. Tente novamente mais tarde."
}
