package service

import (
	"clean-architecture/internal/domain/entity"
	"clean-architecture/internal/interfaces/http/response"
)

type UserResponse struct {
	response.User
}

type UsersResponse []*UserResponse

func NewUser(user *entity.User) *UserResponse {
	return &UserResponse{
		User: response.User{
			ID:    user.ID,
			Email: user.Email,
		},
	}
}

func (u *UserResponse) Response() *response.User {
	return &u.User
}

func NewUsers(users []*entity.User) UsersResponse {
	res := make(UsersResponse, len(users))
	for i := range users {
		res[i] = NewUser(users[i])
	}
	return res
}

func (us UsersResponse) Response() []*response.User {
	res := make([]*response.User, len(us))
	for i := range us {
		res[i] = us[i].Response()
	}
	return res
}
