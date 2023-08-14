package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vrishikesh/go-webserver/helpers"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User User `json:"user"`
}

func CreateUser(p *CreateUserRequest) (*CreateUserResponse, error) {
	var user User
	user.Id = len(usersDB) + 1
	user.Name = p.Name
	usersDB = append(usersDB, user)
	return &CreateUserResponse{User: user}, nil
}

func ParseCreateUser(data []byte) (*CreateUserRequest, error) {
	var req CreateUserRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("could not unmarshal body %s into %T", string(data), req)
		return nil, fmt.Errorf("could not unmarshal body %s into %T", string(data), req)
	}
	return &req, nil
}

func HandleCreateUser(data []byte) *helpers.JsonResponse {
	req, err := ParseCreateUser(data)
	if err != nil {
		return helpers.ErrorResponse(http.StatusBadRequest, err)
	}
	user, err := CreateUser(req)
	if err != nil {
		return helpers.ErrorResponse(http.StatusInternalServerError, err)
	}
	return helpers.SuccessResponse(http.StatusCreated, user)
}
