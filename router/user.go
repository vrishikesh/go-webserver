package router

import (
	"io"
	"net/http"
	"regexp"

	"github.com/vrishikesh/go-webserver/handlers"
	"github.com/vrishikesh/go-webserver/helpers"
)

var SingleUserRegex, _ = regexp.Compile(`/users/(\d+)`)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res *helpers.JsonResponse
	data, _ := io.ReadAll(r.Body)
	// path := strings.TrimRight(r.URL.Path, "/")
	path := r.URL.Path

	switch {
	case r.Method == http.MethodGet && path == "/users":
		res = handlers.HandleGetUsers(r.URL.Query())
	case r.Method == http.MethodGet && SingleUserRegex.MatchString(path):
		res = handlers.HandleGetUser(SingleUserRegex, path)
	case r.Method == http.MethodPost && path == "/users":
		res = handlers.HandleCreateUser(data)
	case r.Method == http.MethodPut && SingleUserRegex.MatchString(path):
		res = handlers.HandleUpdateUser(data, SingleUserRegex, path)
	case r.Method == http.MethodDelete && SingleUserRegex.MatchString(path):
		res = handlers.HandleRemoveUser(SingleUserRegex, path)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Send(w)
}
