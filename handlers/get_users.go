package handlers

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
