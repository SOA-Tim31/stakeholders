package repo

import (
	"fmt"
	"stakeholders/model"
	"strconv"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DatabaseConnection *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{
		DatabaseConnection: db,
	}
}

func (repo *PersonRepository) FindById(id string) (model.Person, error) {
	person := model.Person{}
	dbResult := repo.DatabaseConnection.First(&person, "id = ?", id)
	if dbResult != nil {
		return person, dbResult.Error
	}
	return person, nil
}

func (repo *PersonRepository) FindEmailById(id int) (string, error) {

	person := model.Person{}
	person, err := repo.FindById(strconv.Itoa(id))
	if err != nil {
		return " ", err
	}

	return person.Email, nil
}

func (repo *PersonRepository) CreatePerson(person *model.Person) error {
	dbResult := repo.DatabaseConnection.Create(person)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PersonRepository) Update(person *model.Person) (model.Person, error) {
	var existingPerson model.Person
	dbResult := repo.DatabaseConnection.First(&existingPerson, "id = ?", strconv.Itoa(person.Id))

	if dbResult.Error != nil {
		return model.Person{}, dbResult.Error
	}

	existingPerson.Name = person.Name
	existingPerson.Surname = person.Surname
	existingPerson.Email = person.Email
	existingPerson.ProfileImage = person.ProfileImage
	existingPerson.Bio = person.Bio
	existingPerson.Quote = person.Quote

	if err := repo.DatabaseConnection.Save(&existingPerson).Error; err != nil {
		return model.Person{}, fmt.Errorf("failed to update person: %v", err)
	}

	return existingPerson, nil
}
