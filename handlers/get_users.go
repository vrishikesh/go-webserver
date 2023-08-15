package handlers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/vrishikesh/go-webserver/helpers"
)

var usersDB []User

type GetUsersRequest struct {
	Search string
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(p *GetUsersRequest) (*GetUsersResponse, error) {
	users := make([]User, 0)
	for _, u := range usersDB {
		if strings.Contains(u.Name, p.Search) {
			users = append(users, u)
		}
	}
	return &GetUsersResponse{Users: users}, nil
}

func ParseGetUsers(values url.Values) (*GetUsersRequest, error) {
	var req GetUsersRequest
	if v, ok := values["s"]; ok {
		if len(v) > 0 {
			req.Search = v[0]
		}
	}
	return &req, nil
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	req, err := ParseGetUsers(values)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusBadRequest).Send(w)
		return
	}
	users, err := GetUsers(req)
	if err != nil {
		helpers.NewErrorResponse(err).Send(w)
		return
	}
	helpers.NewSuccessResponse(users).Send(w)
}
