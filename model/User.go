package model

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
