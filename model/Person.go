package model

type Person struct {
	Id           int    `gorm:"column:id;type:integer" json:"id"`
	UserId       int    `json:"userId"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage,omitempty"`
	Bio          string `json:"bio,omitempty"`
	Quote        string `json:"quote,omitempty"`
}
