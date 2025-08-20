package service

import (
	"clean-architecture/internal/domain/entity"
	"clean-architecture/internal/interfaces/http/response"
)

type AuthResponse struct {
	response.Auth
}

func NewAuth(user *entity.User, token string) *AuthResponse {
	return &AuthResponse{
		Auth: response.Auth{
			ID:    user.ID,
			Email: user.Email,
			Token: token,
		},
	}
}

func (a *AuthResponse) Response() *response.Auth {
	return &a.Auth
}
