package router

import (
	"net/http"
	"regexp"

	"github.com/vrishikesh/go-webserver/handlers"
)

var GetUserRegex, _ = regexp.Compile(`/users/(\d+)`)

func UserRouter(r *http.Request, data []byte) (any, error) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		req, err := handlers.ParseGetUsers(data)
		if err != nil {
			return nil, err
		}
		return handlers.GetUsers(req)
	case r.Method == http.MethodGet && GetUserRegex.MatchString(r.URL.Path):
		req, err := handlers.ParseGetUser(GetUserRegex, r.URL.Path)
		if err != nil {
			return nil, err
		}
		return handlers.GetUser(req)
	default:
		return handlers.MethodNotAllowed(nil)
	}
}
