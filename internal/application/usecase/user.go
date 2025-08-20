//go:generate mockgen -source=$GOFILE -package mock_$GOPACKAGE -destination=./../../../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"clean-architecture/internal/domain/dto"
	"clean-architecture/internal/domain/entity"
	"clean-architecture/internal/domain/repository"
	"context"
)

type userUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepository: userRepository}
}

type UserUsecase interface {
	List(ctx context.Context, params *dto.UserListParams) ([]*entity.User, error)
	Get(ctx context.Context, id uint64) (*entity.User, error)
}

func (uc *userUsecase) List(ctx context.Context, params *dto.UserListParams) ([]*entity.User, error) {
	users, err := uc.UserRepository.List(ctx, params)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *userUsecase) Get(ctx context.Context, id uint64) (*entity.User, error) {
	user, err := uc.UserRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
