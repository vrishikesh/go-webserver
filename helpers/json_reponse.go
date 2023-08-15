package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Status  int    `json:"-"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(data any, status ...int) *JsonResponse {
	currentStatus := http.StatusOK
	if len(status) > 0 {
		currentStatus = status[0]
	}

	return &JsonResponse{
		Status:  currentStatus,
		Success: true,
		Error:   "",
		Data:    data,
	}
}

func NewErrorResponse(err error, status ...int) *JsonResponse {
	currentStatus := http.StatusInternalServerError
	if len(status) > 0 {
		currentStatus = status[0]
	}

	return &JsonResponse{
		Status:  currentStatus,
		Success: false,
		Error:   err.Error(),
	}
}

func (res *JsonResponse) Send(w http.ResponseWriter) {
	if res.Error != "" {
		log.Printf("error: %s", res.Error)
	}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal response: %s", err)
		if _, err := w.Write([]byte("something went wrong")); err != nil {
			log.Printf("could not write to stdout: %s", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if _, err := w.Write(b); err != nil {
		log.Printf("could not write to stdout: %s", err)
	}
}
