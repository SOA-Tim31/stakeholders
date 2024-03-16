package repo

import (
	"stakeholders/model"

	"gorm.io/gorm"
)

type AppRatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AppRatingRepository) FindAll() ([]model.AppRating, error) {
	var appRatings []model.AppRating
	dbResult := repo.DatabaseConnection.Find(&appRatings)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return appRatings, nil
}

func (repo *AppRatingRepository) Create(appRating *model.AppRating) (*model.AppRating, error) {
	dbResult := repo.DatabaseConnection.Create(appRating)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return appRating, nil
}
