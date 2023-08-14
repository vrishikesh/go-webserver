package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vrishikesh/go-webserver/helpers"
)

var usersDB []User

type GetUsersRequest struct {
	Search string `json:"search"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(p *GetUsersRequest) (*GetUsersResponse, error) {
	users := make([]User, len(usersDB))
	copy(users, usersDB)
	return &GetUsersResponse{Users: users}, nil
}

func ParseGetUsers(data []byte) (*GetUsersRequest, error) {
	var req GetUsersRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("could not unmarshal body %s into %T", string(data), req)
		return nil, fmt.Errorf("could not unmarshal body %s into %T", string(data), req)
	}
	return &req, nil
}

func HandleGetUsers(data []byte) *helpers.JsonResponse {
	req, err := ParseGetUsers(data)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err)
	}
	users, err := GetUsers(req)
	if err != nil {
		return helpers.ErrorResponse(http.StatusInternalServerError, err)
	}
	return helpers.SuccessResponse(http.StatusOK, users)
}
