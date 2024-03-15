package handler

import (
	"encoding/json"
	"net/http"
	"stakeholders/dto"
	"stakeholders/model"
	"stakeholders/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) FindAllUsers(writer http.ResponseWriter, req *http.Request) {
	users, err := handler.UserService.FindAllUsers()
	println("username", users[0].Username)
	println("id:", users[0].Id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var accountDtos []dto.AccountDto = handler.MapUsersToAccounts(users)

	json.NewEncoder(writer).Encode(accountDtos)
}

func (handler *UserHandler) MapUsersToAccounts(users []model.User) []dto.AccountDto {

	var accountDtos []dto.AccountDto
	for _, user := range users {
		accountDto := dto.AccountDto{
			UserId:   user.Id,
			Username: user.Username,
			Password: user.Password,
			Email:    handler.UserService.FindEmail(&user), // You might need to fill this from somewhere
			Role:     MapToString(user.Role),
			IsActive: user.IsActive,
		}
		accountDtos = append(accountDtos, accountDto)
	}
	return accountDtos
}

func MapAccountToUser(accountDto dto.AccountDto) model.User {
	user := model.User{
		Id:       accountDto.UserId,
		Username: accountDto.Username,
		Password: accountDto.Password,
		Role:     MapToUserRole(accountDto.Role),
		IsActive: accountDto.IsActive,
	}
	return user
}

func MapToString(role model.UserRole) string {
	switch role {
	case model.Administrator:
		return "Administrator"
	case model.Author:
		return "Author"
	case model.Tourist:
		return "Tourist"
	default:
		return "Tourist"
	}
}

func MapToUserRole(role string) model.UserRole {
	switch role {
	case "Administrator":
		return model.Administrator
	case "Author":
		return model.Author
	case "Tourist":
		return model.Tourist
	default:
		return model.Tourist
	}
}

func (handler *UserHandler) BlockOrUblock(writer http.ResponseWriter, req *http.Request) {
	var accountDto dto.AccountDto
	err := json.NewDecoder(req.Body).Decode(&accountDto)
	user := MapAccountToUser(accountDto)
	println("prvi")
	if err != nil {
		println("Error while parsing ", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUser, err2 := handler.UserService.BlockOrUblock(&user)
	if err2 != nil {
		println("Error while updateing a new user")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedUser)

}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {

}
