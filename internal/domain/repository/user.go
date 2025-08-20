//go:generate mockgen -source=$GOFILE -package mock_$GOPACKAGE -destination=./../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"clean-architecture/internal/application/dto"
	"clean-architecture/internal/domain/entity"
	"context"
)

type UserRepository interface {
	List(ctx context.Context, params *dto.UserListParams) ([]*entity.User, error)
	Get(ctx context.Context, userID uint64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
