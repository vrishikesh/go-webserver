package handlers

type MethodNotAllowedRequest struct{}
type MethodNotAllowedResponse struct {
	Error string `json:"error"`
}

func MethodNotAllowed(*MethodNotAllowedRequest) (*MethodNotAllowedResponse, error) {
	return &MethodNotAllowedResponse{
		Error: "method not allowed",
	}, nil
}
