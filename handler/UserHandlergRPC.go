package handler

import (
	"context"
	"fmt"
	"stakeholders/model"
	stakeholders "stakeholders/proto"
	"stakeholders/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandlergRPC struct {
	UserService *service.UserService
	AuthService *service.AuthService
	stakeholders.UnimplementedStakeholderServiceServer
}

func NewUserHandlergRPC(us *service.UserService, as *service.AuthService) *UserHandlergRPC {
	return &UserHandlergRPC{
		UserService: us,
		AuthService: as,
	}
}

func (uh *UserHandlergRPC) RegistrationRpc(ctx context.Context, req *stakeholders.RegistrationRequest) (*stakeholders.RegistrationResponse, error) {
	registration := model.Registration{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
		Surname:  req.Surname,
		Role:     req.Role,
	}

	token := uh.AuthService.GenerateUniqueVerificationToken()
	item := false

	err := uh.UserService.Registration(&registration, &token, &item)
	if err != nil {
		return nil, fmt.Errorf("error while registering a new user: %v", err)
	}

	err = uh.AuthService.SendVerificationMail(&registration, token)
	if err != nil {
		return nil, fmt.Errorf("Error whie sending a email: %v", err)
	}

	return &stakeholders.RegistrationResponse{
		Message: "User registered successfully",
	}, nil

}

func (uh *UserHandlergRPC) GetProfileRpc(ctx context.Context, req *stakeholders.GetProfileRequest) (*stakeholders.GetProfileResponse, error) {
	userId := req.Id
	person, err := uh.UserService.GetPersonByUserId(&userId)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	return &stakeholders.GetProfileResponse{
		Person: &stakeholders.Person{
			Id:           uint64(person.Id),
			UserId:       uint64(person.UserId),
			Name:         person.Name,
			Surname:      person.Surname,
			ProfileImage: person.ProfileImage,
			Email:        person.Email,
			Bio:          person.Bio,
			Quote:        person.Quote,
		},
	}, nil
}

func (uh *UserHandlergRPC) FindAllUsersRpc(ctx context.Context, req *stakeholders.FindAllRequest) (*stakeholders.FindAllResponse, error) {
	users, err := uh.UserService.FindAllUsers()

	if err != nil {
		return nil, fmt.Errorf("error while finding all users: %v", err)
	}

	var accountDtos []*stakeholders.AccountDto = uh.MapUsersToAccounts(users)

	return &stakeholders.FindAllResponse{
		Accounts: accountDtos,
	}, nil
}

func (uh *UserHandlergRPC) MapUsersToAccounts(users []model.User) []*stakeholders.AccountDto {

	var accountDtos []*stakeholders.AccountDto
	for _, user := range users {
		accountDto := stakeholders.AccountDto{
			UserId:   uint64(user.Id),
			Username: user.Username,
			Password: user.Password,
			Email:    uh.UserService.FindEmail(&user),
			Role:     MapToString(user.Role),
			IsActive: user.IsActive,
		}
		accountDtos = append(accountDtos, &accountDto)
	}
	return accountDtos
}
