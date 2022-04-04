package model

import "time"

type User struct {
	UserId   string
	FullName string
	Email    string
	PassWord string
	Role     string
	CreateAt time.Time
	UpdateAt time.Time
	Token    string
}
