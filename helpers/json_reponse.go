package helpers

type JsonResponse struct {
	Status  int   `json:"-"`
	Success bool  `json:"success"`
	Error   error `json:"error"`
	Data    any   `json:"data,omitempty"`
}

func SuccessResponse(status int, data any) *JsonResponse {
	return &JsonResponse{
		Status:  status,
		Success: true,
		Error:   nil,
		Data:    data,
	}
}

func ErrorResponse(status int, err error) *JsonResponse {
	return &JsonResponse{
		Status:  status,
		Success: false,
		Error:   err,
	}
}
