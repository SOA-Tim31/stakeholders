package repo

import (
	"stakeholders/model"

	"gorm.io/gorm"
)

type PersonRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PersonRepository) FindById(id string) (model.Person, error) {
	person := model.Person{}
	dbResult := repo.DatabaseConnection.First(&person, "id = ?", id)
	if dbResult != nil {
		return person, dbResult.Error
	}
	return person, nil
}

func (repo *PersonRepository) CreatePerson(person *model.Person) error {
	dbResult := repo.DatabaseConnection.Create(person)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
