package handlers

type ResourceNotFoundRequest struct{}
type ResourceNotFoundResponse struct {
	Error string `json:"error"`
}

func ResourceNotFound(*ResourceNotFoundRequest) (*ResourceNotFoundResponse, error) {
	return &ResourceNotFoundResponse{
		Error: "resource not found",
	}, nil
}
