package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vrishikesh/go-webserver/helpers"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	User User `json:"user"`
}

func CreateUser(p *CreateUserRequest) (*CreateUserResponse, error) {
	if strings.TrimSpace(p.Name) == "" {
		log.Printf("name cannot be empty")
		return nil, fmt.Errorf("name cannot be empty")
	}

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
		return nil, fmt.Errorf("could not parse body")
	}
	return &req, nil
}

func HandleCreateUser(data []byte) *helpers.JsonResponse {
	req, err := ParseCreateUser(data)
	if err != nil {
		return helpers.NewErrorResponse(err, http.StatusBadRequest)
	}
	user, err := CreateUser(req)
	if err != nil {
		return helpers.NewErrorResponse(err, http.StatusInternalServerError)
	}
	return helpers.NewSuccessResponse(user, http.StatusCreated)
}
