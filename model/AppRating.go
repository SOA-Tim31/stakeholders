package model

import "time"

type AppRating struct {
	Id          int       `gorm:"column:id;type:integer" json:"id"`
	UserId      int       `json:"userId"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
}
