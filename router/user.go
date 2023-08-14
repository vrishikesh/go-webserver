package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/vrishikesh/go-webserver/handlers"
)

var GetUserRegex, _ = regexp.Compile(`/users/(\d+)`)

func UserRouter(r *http.Request, data []byte) (any, error) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/users":
		var req handlers.GetUsersRequest
		err := json.Unmarshal(data, &req)
		if err != nil {
			log.Printf("could not unmarshal body %s into %T", string(data), req)
			return nil, fmt.Errorf("could not unmarshal body %s into %T", string(data), req)
		}
		return handlers.GetUsers(&req)
	case r.Method == http.MethodGet && GetUserRegex.MatchString(r.URL.Path):
		var req handlers.GetUserRequest
		sss := GetUserRegex.FindAllStringSubmatch(r.URL.Path, -1)
		ss := sss[0]
		if len(ss) < 2 {
			log.Printf("could not find param user id: %v", ss)
			return nil, fmt.Errorf("could not find param user id%v", ss)
		}
		req.Id, _ = strconv.Atoi(ss[1])
		return handlers.GetUser(&req)
	default:
		var req handlers.MethodNotAllowedRequest
		return handlers.MethodNotAllowed(&req)
	}
}
