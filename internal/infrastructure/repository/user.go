package repository

import (
	"clean-architecture/internal/application/dto"
	"clean-architecture/internal/domain/entity"
	"clean-architecture/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{DB: db}
}

func (ur *userRepo) List(ctx context.Context, params *dto.UserListParams) ([]*entity.User, error) {
	var users entity.Users
	if err := ur.DB.WithContext(ctx).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users).Error; err != nil {
		return nil, dbError(err)
	}
	return users, nil
}

func (ur *userRepo) Get(ctx context.Context, id uint64) (*entity.User, error) {
	var user *entity.User
	if err := ur.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, dbError(err)
	}
	return user, nil
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	if err := ur.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, dbError(err)
	}
	return user, nil
}
