package handlers

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
