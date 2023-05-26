package utils

type HandlerError struct {
	error
	Code int
	Message string
}
func ValidateCpf(cpf string) *HandlerError {
	if len(cpf) == 11 {
		return nil
	}
	return &HandlerError{
		Code: 400,
		Message: "Cpf inválido",
	}
}
