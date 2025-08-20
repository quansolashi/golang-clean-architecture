package usecase

import (
	"context"
	"time"

	"clean-architecture/internal/domain/entity"
	derror "clean-architecture/internal/domain/error"
	"clean-architecture/internal/domain/repository"
	"clean-architecture/internal/domain/service/auth"
	"clean-architecture/internal/infrastructure/auth/security"
)

type authUsecase struct {
	UserRepository repository.UserRepository
	TokenService   auth.TokenService
}

func NewAuthUsecase(
	userRepository repository.UserRepository,
	tokenService auth.TokenService,
) AuthUsecase {
	return &authUsecase{
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
}

type AuthUsecase interface {
	Authenticate(ctx context.Context, email, password string) (*entity.User, string, error)
}

func (uc *authUsecase) Authenticate(ctx context.Context, email, password string) (*entity.User, string, error) {
	user, err := uc.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", derror.Unauthorized("invalid email")
	}

	if err := security.VerifyPassword(user.Password, password); err != nil {
		return nil, "", derror.Unauthorized("invalid password")
	}

	token, err := uc.TokenService.GenerateToken(ctx, user.ID, time.Hour*24)
	if err != nil {
		return nil, "", derror.InternalServer(err)
	}

	return user, token, nil
}
