package usecase

import (
	"clean-architecture/internal/domain/repository"
	"clean-architecture/internal/domain/service/auth"
)

type Params struct {
	Repository   *repository.Repository
	TokenService auth.TokenService
}

type Usecase struct {
	User UserUsecase
	Auth AuthUsecase
}

func NewUsecase(params *Params) *Usecase {
	return &Usecase{
		User: NewUserUsecase(params.Repository.User),
		Auth: NewAuthUsecase(params.Repository.User, params.TokenService),
	}
}
