package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/vrishikesh/go-webserver/helpers"
)

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

func GetUser(p *GetUserRequest) (*GetUserResponse, error) {
	var user User
	for _, u := range usersDB {
		if u.Id == p.Id {
			user = u
			break
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("could not find user with id %d", p.Id)
	}
	return &GetUserResponse{User: user}, nil
}

func ParseGetUser(regex *regexp.Regexp, path string) (*GetUserRequest, error) {
	var req GetUserRequest
	sss := regex.FindAllStringSubmatch(path, -1)
	ss := sss[0]
	if len(ss) < 2 {
		log.Printf("could not find param user id: %v", ss)
		return nil, fmt.Errorf("could not find param user id: %v", ss)
	}
	req.Id, _ = strconv.Atoi(ss[1])
	return &req, nil
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	req, err := ParseGetUser(helpers.UserRouteRegex, r.URL.Path)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusBadRequest).Send(w)
		return
	}
	data, err := GetUser(req)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusInternalServerError).Send(w)
		return
	}
	helpers.NewSuccessResponse(data, http.StatusOK).Send(w)
}
