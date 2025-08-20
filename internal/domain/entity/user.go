package entity

import "time"

type User struct {
	ID        uint64
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []*User
