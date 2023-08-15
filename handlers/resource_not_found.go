package handlers

import (
	"net/http"

	"github.com/vrishikesh/go-webserver/helpers"
)

func HandleResourceNotFound() *helpers.JsonResponse {
	err := helpers.NewResourceNotFoundError()
	return helpers.NewErrorResponse(err, http.StatusNotFound)
}

func HandleHandleResourceNotFoundRoute(w http.ResponseWriter, r *http.Request) {
	res := HandleResourceNotFound()
	res.Send(w)
}
