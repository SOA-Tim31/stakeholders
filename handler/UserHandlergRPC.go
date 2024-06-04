package handler

import (
	"context"
	"fmt"
	"stakeholders/model"
	stakeholders "stakeholders/proto"
	"stakeholders/service"
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

	err, idUser:= uh.UserService.Registration(&registration, &token, &item)
	if err != nil {
		return nil, fmt.Errorf("error while registering a new user: %v", err)
	}

	err = uh.AuthService.SendVerificationMail(&registration, token)
	if err != nil {
		return nil, fmt.Errorf("Error whie sending a email: %v", err)
	}

	return &stakeholders.RegistrationResponse{
		Message: "User registered successfully",
		Id: uint64(idUser),
	}, nil

}
