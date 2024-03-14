package service

import (
	"fmt"
	"stakeholders/model"
	"stakeholders/repo"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

func (service *UserService) FindUser(id string) (*model.User, error) {
	user, err := service.UserRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &user, nil
}

func (service *UserService) Create(user *model.User) error {
	err := service.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) BlockOrUblock(user *model.User) (*model.User, error) {
	updatedUser, err := service.UserRepo.BlockOrUblock(user)
	if err != nil {
		return nil, fmt.Errorf("couldn't do block/unblock operation")
	}
	return &updatedUser, nil
}

func (service *UserService) FindAllUsers() ([]model.User, error) {
	return service.UserRepo.FindAll()
}
