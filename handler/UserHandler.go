package handler

import (
	"encoding/json"
	"net/http"
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
	json.NewEncoder(writer).Encode(users)
}

func MapTempUserToUser(tempUser model.TempUser) model.User {
	user := model.User{
		Id:       tempUser.UserId,
		Username: tempUser.Username,
		Password: tempUser.Password,
		Email:    tempUser.Email,
		Role:     MapToUserRole(tempUser.RoleString),
		IsActive: tempUser.IsActive,
	}
	return user
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
	var tempUser model.TempUser
	err := json.NewDecoder(req.Body).Decode(&tempUser)
	user := MapTempUserToUser(tempUser)

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
