package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/vrishikesh/go-webserver/helpers"
)

type UpdateUserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateUserResponse struct {
	User User `json:"user"`
}

func UpdateUser(p *UpdateUserRequest) (*UpdateUserResponse, error) {
	var user User
	for i := range usersDB {
		u := &usersDB[i]
		if u.Id == p.Id {
			u.Name = p.Name
			user = *u
			break
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("could not update user with id %d", p.Id)
	}
	return &UpdateUserResponse{User: user}, nil
}

func ParseUpdateUser(data []byte, regex *regexp.Regexp, path string) (*UpdateUserRequest, error) {
	var req UpdateUserRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("could not unmarshal body %s into %T", string(data), req)
		return nil, fmt.Errorf("could not unmarshal body %s into %T", string(data), req)
	}

	sss := regex.FindAllStringSubmatch(path, -1)
	ss := sss[0]
	if len(ss) < 2 {
		log.Printf("could not find param user id: %v", ss)
		return nil, fmt.Errorf("could not find param user id: %v", ss)
	}
	req.Id, _ = strconv.Atoi(ss[1])

	return &req, nil
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	req, err := ParseUpdateUser(b, helpers.UserRouteRegex, r.URL.Path)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusBadRequest).Send(w)
		return
	}
	user, err := UpdateUser(req)
	if err != nil {
		helpers.NewErrorResponse(err).Send(w)
		return
	}
	helpers.NewSuccessResponse(user).Send(w)
}
