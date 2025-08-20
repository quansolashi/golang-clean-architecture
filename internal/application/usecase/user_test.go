package usecase

import (
	"clean-architecture/internal/domain/dto"
	"clean-architecture/internal/domain/entity"
	mock_repository "clean-architecture/mock/repository"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserUsecase_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userUsecase := NewUserUsecase(userRepo)

	users := []*entity.User{
		{
			ID:    1,
			Email: "user1@example.com",
		},
		{
			ID:    2,
			Email: "user2@example.com",
		},
		{
			ID:    3,
			Email: "user3@example.com",
		},
	}

	type args struct {
		param *dto.UserListParams
	}
	type want struct {
		users  []*entity.User
		hasErr bool
	}
	tests := []struct {
		name  string
		setup func(ctx context.Context, t *testing.T)
		args  args
		want  want
	}{
		{
			name: "Success",
			setup: func(ctx context.Context, t *testing.T) {
				userRepo.EXPECT().
					List(ctx, gomock.Any()).
					Return(users[1:], nil)
			},
			args: args{
				param: &dto.UserListParams{
					Limit:  2,
					Offset: 1,
				},
			},
			want: want{
				users:  users[1:],
				hasErr: false,
			},
		},
		{
			name: "Repository Error",
			setup: func(ctx context.Context, t *testing.T) {
				userRepo.EXPECT().
					List(ctx, gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			args: args{
				param: &dto.UserListParams{Limit: 10, Offset: 0},
			},
			want: want{
				users:  nil,
				hasErr: true,
			},
		},
		{
			name: "Empty Result",
			setup: func(ctx context.Context, t *testing.T) {
				userRepo.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*entity.User{}, nil)
			},
			args: args{
				param: &dto.UserListParams{Limit: 10, Offset: 0},
			},
			want: want{
				users:  []*entity.User{},
				hasErr: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := t.Context()
			tt.setup(ctx, t)

			actual, err := userUsecase.List(ctx, tt.args.param)
			assert.Equal(t, tt.want.hasErr, err != nil, err)
			assert.Equal(t, tt.want.users, actual)
		})
	}
}
