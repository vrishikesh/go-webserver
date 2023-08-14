package router

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

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
	path := strings.TrimRight(r.URL.Path, "/")

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

	b, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Status)
	if res.Error != "" {
		log.Printf("error: %s", res.Error)
	}
	if _, err := w.Write(b); err != nil {
		log.Printf("could not write to stdout: %s", err)
	}
}

func UserRouter2(r *http.Request, data []byte) *helpers.JsonResponse {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		return handlers.HandleGetUsers(r.URL.Query())
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
