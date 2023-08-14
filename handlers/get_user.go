package handlers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

func GetUser(p *GetUserRequest) (*GetUserResponse, error) {
	return &GetUserResponse{
		User: User{
			Id:   p.Id,
			Name: "random",
		},
	}, nil
}

func ParseGetUser(regex *regexp.Regexp, path string) (*GetUserRequest, error) {
	var req GetUserRequest
	sss := regex.FindAllStringSubmatch(path, -1)
	ss := sss[0]
	if len(ss) < 2 {
		log.Printf("could not find param user id: %v", ss)
		return nil, fmt.Errorf("could not find param user id%v", ss)
	}
	req.Id, _ = strconv.Atoi(ss[1])
	return &req, nil
}
