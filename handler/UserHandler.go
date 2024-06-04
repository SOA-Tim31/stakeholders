package handler

import (
	"encoding/json"
	"net/http"
	"stakeholders/dto"
	"stakeholders/model"
	"stakeholders/service"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService *service.UserService
	AuthService *service.AuthService
}

func (handler *UserHandler) FindAllUsers(writer http.ResponseWriter, req *http.Request) {
	users, err := handler.UserService.FindAllUsers()

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

// func (userHandler *UserHandler) Registration(writer http.ResponseWriter, req *http.Request) {
// 	var registration model.Registration

// 	err := json.NewDecoder(req.Body).Decode(&registration)
// 	if err != nil {
// 		println("Error while parsing json")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	token := userHandler.AuthService.GenerateUniqueVerificationToken()
// 	item := false

// 	err = userHandler.UserService.Registration(&registration, &token, &item)
// 	if err != nil {
// 		println("Error while registering a new user")
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}

// 	err = userHandler.AuthService.SendVerificationMail(&registration, token)
// 	if err != nil {
// 		println("Error while sending an email")
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusCreated)
// 	writer.Header().Set("Content-Type", "application/json")
// }

func (userHandler *UserHandler) VerifyEmail(writer http.ResponseWriter, req *http.Request) {
	token := mux.Vars(req)["token"]
	user, err := userHandler.UserService.GetAndVerifyUserByToken(&token)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)

}

/* bez tokena
func (handler *UserHandler) Register(writer http.ResponseWriter, req *http.Request) {

	var accountRegDto dto.AccountRegistrationDto
	err := json.NewDecoder(req.Body).Decode(&accountRegDto)

	if err != nil {
		println("Error while parsing ", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	authTokenDto, err := handler.UserService.Register(&accountRegDto)

	if err != nil {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(authTokenDto)
}*/
