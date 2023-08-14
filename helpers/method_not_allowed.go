package helpers

type MethodNotAllowedError struct{}

func (m MethodNotAllowedError) Error() string {
	return "method not allowed"
}

func NewMethodNotAllowedError() *MethodNotAllowedError {
	return &MethodNotAllowedError{}
}
