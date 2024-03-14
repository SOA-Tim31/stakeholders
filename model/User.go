package model

import (
	"gorm.io/gorm"
)

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
	IsActive bool     `json:"isActive"`
}

type TempUser struct {
	UserId     int
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	RoleString string `json:"role"`
	IsActive   bool   `json:"isActive"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	//user.Id = uuid.New()
	return nil
}
