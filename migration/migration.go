package migration

import (
	"stakeholders/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Student{}, &model.User{}, &model.Person{}); err != nil {
		return err
	}
	return nil
}
