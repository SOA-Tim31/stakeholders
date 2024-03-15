package repo

import (
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
	println("name", person.Name)
	println("email", person.Email)
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
