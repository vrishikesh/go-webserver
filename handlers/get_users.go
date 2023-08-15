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

func HandleGetUsers(values url.Values) *helpers.JsonResponse {
	req, err := ParseGetUsers(values)
	if err != nil {
		return helpers.NewErrorResponse(err, http.StatusBadRequest)
	}
	users, err := GetUsers(req)
	if err != nil {
		return helpers.NewErrorResponse(err, http.StatusInternalServerError)
	}
	return helpers.NewSuccessResponse(users, http.StatusOK)
}
