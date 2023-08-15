package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/vrishikesh/go-webserver/helpers"
)

type RemoveUserRequest struct {
	Id int `json:"id"`
}

type RemoveUserResponse struct{}

func RemoveUser(p *RemoveUserRequest) (*RemoveUserResponse, error) {
	var user User
	for i, u := range usersDB {
		if u.Id == p.Id {
			user = u
			usersDB = append(usersDB[:i], usersDB[i+1:]...)
			break
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("could not remove user")
	}
	return nil, nil
}

func ParseRemoveUser(regex *regexp.Regexp, path string) (*RemoveUserRequest, error) {
	var req RemoveUserRequest
	sss := regex.FindAllStringSubmatch(path, -1)
	ss := sss[0]
	if len(ss) < 2 {
		log.Printf("could not find param user id: %v", ss)
		return nil, fmt.Errorf("could not find param user id: %v", ss)
	}
	req.Id, _ = strconv.Atoi(ss[1])
	return &req, nil
}

func HandleRemoveUser(w http.ResponseWriter, r *http.Request) {
	req, err := ParseRemoveUser(helpers.UserRouteRegex, r.URL.Path)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusBadRequest).Send(w)
		return
	}
	data, err := RemoveUser(req)
	if err != nil {
		helpers.NewErrorResponse(err, http.StatusInternalServerError).Send(w)
		return
	}
	helpers.NewSuccessResponse(data, http.StatusNoContent).Send(w)
}
