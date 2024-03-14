package service

import (
	"fmt"
	"stakeholders/model"
	"stakeholders/repo"
)

type PersonService struct {
	PersonRepo *repo.PersonRepository
}

func (service *PersonService) FindPerson(id string) (*model.Person, error) {
	person, err := service.PersonRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &person, nil
}

func (service *PersonService) Create(person *model.Person) error {
	err := service.PersonRepo.CreatePerson(person)
	if err != nil {
		return err
	}
	return nil
}
