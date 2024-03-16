package repo

import (
	"stakeholders/model"
	"strconv"

	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) FindById(id string) (model.User, error) {
	user := model.User{}
	dbResult := repo.DatabaseConnection.First(&user, "id = ?", id)
	if dbResult != nil {
		return user, dbResult.Error
	}
	return user, nil
}

func (repo *UserRepository) BlockOrUblock(user *model.User) (model.User, error) {
	var foundUser model.User
	dbResult := repo.DatabaseConnection.First(&foundUser, "id = ?", strconv.Itoa(user.Id))

	if dbResult.Error != nil {
		return model.User{}, dbResult.Error
	}

	foundUser.IsActive = !foundUser.IsActive
	updateResult := repo.DatabaseConnection.Save(&foundUser)

	if updateResult.Error != nil {
		return model.User{}, updateResult.Error
	}

	return foundUser, nil
}

func (repo *UserRepository) CreateUser(user *model.User) (int, error) {
	dbResult := repo.DatabaseConnection.Create(user)
	if dbResult.Error != nil {
		return 0, dbResult.Error
	}
	return user.Id, nil
}

func (repo *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	dbResult := repo.DatabaseConnection.Find(&users)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return users, nil
}
