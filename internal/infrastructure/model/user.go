package model

import (
	"clean-architecture/internal/domain/entity"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;<-:create"`
	Email     string    `gorm:"not null;size:256;unique"`
	Password  string    `gorm:"not null;size:256"`
	CreatedAt time.Time `gorm:"not null;<-:create"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Users []*User

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (us Users) ToEntities() entity.Users {
	entities := make(entity.Users, len(us))
	for i, user := range us {
		entities[i] = user.ToEntity()
	}
	return entities
}
