package handlers

import (
	"encoding/json"
	"fmt"
	"log"
)

type GetUsersRequest struct {
	Id []int `json:"id"`
}

type GetUsersResponse struct {
	Body GetUsersResponseBody `json:"body"`
}

type GetUsersResponseBody struct {
	Users []User `json:"users"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(p *GetUsersRequest) (*GetUsersResponse, error) {
	users := []User{}
	for _, v := range p.Id {
		user := User{Id: v, Name: "random"}
		users = append(users, user)
	}
	responseBody := GetUsersResponseBody{Users: users}
	return &GetUsersResponse{Body: responseBody}, nil
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
