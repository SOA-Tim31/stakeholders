package model

import (
	"github.com/golang-jwt/jwt"
)

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

type User struct {
	Id                int      `gorm:"column:id;type:integer" json:"id"`
	Username          string   `json:"username"`
	Password          string   `json:"password"`
	Role              UserRole `json:"role"`
	IsActive          bool     `json:"isActive"`
	VerificationToken string   `json:"verificationToken"`
}

type TempUser struct {
	Id                int
	Username          string `json:"username"`
	Password          string `json:"password"`
	RoleString        string `json:"role"`
	IsActive          bool   `json:"isActive"`
	VerificationToken string `json:"verificationToken"`
}

func (u *User) GetRoleName() string {
	switch u.Role {
	case Administrator:
		return "administrator"
	case Tourist:
		return "tourist"
	case Author:
		return "author"
	default:
		return ""
	}
}

func ParseUserRole(role string) UserRole {
	switch role {
	case "administrator":
		return Administrator
	case "tourist":
		return Tourist
	case "author":
		return Author
	default:
		return Tourist // Defaultna vrijednost, možete promijeniti prema potrebi
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
	jwt.StandardClaims
}

type Registration struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Role     string `json:"Role"`
}
