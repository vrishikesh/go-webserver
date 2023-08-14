package helpers

type ResourceNotFoundError struct{}

func (m ResourceNotFoundError) Error() string {
	return "resource not found"
}

func NewResourceNotFoundError() *ResourceNotFoundError {
	return &ResourceNotFoundError{}
}
