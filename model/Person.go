package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" gorm:"not null;type:string"`
}

func (person *Person) BeforeCreate(scope *gorm.DB) error {
	person.ID = uuid.New()
	return nil
}
