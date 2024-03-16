package service

import (
	"fmt"
	"stakeholders/model"
	"stakeholders/repo"
)

type AppRatingService struct {
	AppRatingRepository *repo.AppRatingRepository
}

func (service *AppRatingService) GetAll() ([]model.AppRating, error) {
	return service.AppRatingRepository.FindAll()
}

func (service *AppRatingService) Create(appRating *model.AppRating) (*model.AppRating, error) {
	appRatings, err := service.AppRatingRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't find all ratings")
	}
	if len(appRatings) != 0 {
		for _, rate := range appRatings {

			if rate.UserId == appRating.UserId {
				return nil, fmt.Errorf("this user has already rated")
			}
		}
	}
	createdRating, err2 := service.AppRatingRepository.Create(appRating)

	if err2 != nil {
		return nil, fmt.Errorf("couldn't create")
	}

	return createdRating, nil
}
