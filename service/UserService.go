package service

import (
	"fmt"
	"stakeholders/model"
	"stakeholders/repo"
)

//var personService = PersonService{PersonRepo: &repo.PersonRepository{}}

type UserService struct {
	UserRepo      *repo.UserRepository
	PersonService *PersonService
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
	println("drugi")
	updatedUser, err := service.UserRepo.BlockOrUblock(user)
	if err != nil {
		return nil, fmt.Errorf("couldn't do block/unblock operation")
	}
	return &updatedUser, nil
}

func (service *UserService) FindEmail(user *model.User) string {
	println("3.5")
	email, err := service.PersonService.FindEmail(user.Id)
	if err != nil {
		return "Invalid email"
	}
	return email
}

func (service *UserService) FindAllUsers() ([]model.User, error) {
	return service.UserRepo.FindAll()
}
