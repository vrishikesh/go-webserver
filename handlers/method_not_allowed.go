package handlers

import (
	"net/http"

	"github.com/vrishikesh/go-webserver/helpers"
)

func HandleMethodNotAllowed() *helpers.JsonResponse {
	err := helpers.NewMethodNotAllowedError()
	return helpers.ErrorResponse(http.StatusMethodNotAllowed, err)
}
