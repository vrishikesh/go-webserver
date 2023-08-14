package router

import (
	"net/http"
	"regexp"

	"github.com/vrishikesh/go-webserver/handlers"
	"github.com/vrishikesh/go-webserver/helpers"
)

var SingleUserRegex, _ = regexp.Compile(`/users/(\d+)`)

func UserRouter(r *http.Request, data []byte) *helpers.JsonResponse {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		return handlers.HandleGetUsers(data)
	case r.Method == http.MethodGet && SingleUserRegex.MatchString(r.URL.Path):
		return handlers.HandleGetUser(SingleUserRegex, r.URL.Path)
	case r.Method == http.MethodPost && r.URL.Path == "/users":
		return handlers.HandleCreateUser(data)
	case r.Method == http.MethodPut && SingleUserRegex.MatchString(r.URL.Path):
		return handlers.HandleUpdateUser(data, SingleUserRegex, r.URL.Path)
	case r.Method == http.MethodDelete && SingleUserRegex.MatchString(r.URL.Path):
		return handlers.HandleRemoveUser(SingleUserRegex, r.URL.Path)
	default:
		return handlers.HandleMethodNotAllowed()
	}
}
