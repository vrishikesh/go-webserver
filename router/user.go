package router

import (
	"net/http"
	"regexp"

	"github.com/vrishikesh/go-webserver/handlers"
	"github.com/vrishikesh/go-webserver/helpers"
)

var GetPutUserRegex, _ = regexp.Compile(`/users/(\d+)`)

func UserRouter(r *http.Request, data []byte) *helpers.JsonResponse {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		return handlers.HandleGetUsers(data)
	case r.Method == http.MethodGet && GetPutUserRegex.MatchString(r.URL.Path):
		return handlers.HandleGetUser(GetPutUserRegex, r.URL.Path)
	case r.Method == http.MethodPost && r.URL.Path == "/users":
		return handlers.HandleCreateUser(data)
	case r.Method == http.MethodPut && GetPutUserRegex.MatchString(r.URL.Path):
		return handlers.HandleUpdateUser(data, GetPutUserRegex, r.URL.Path)
	default:
		return handlers.HandleMethodNotAllowed()
	}
}
