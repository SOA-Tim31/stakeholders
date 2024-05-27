package service

import (
	"fmt"
	"stakeholders/model"
	"stakeholders/repo"
)

//var personService = PersonService{PersonRepo: &repo.PersonRepository{}}

type UserService struct {
	UserRepo      *repo.UserRepository
	PersonService *PersonService
}

func (service *UserService) FindUser(id string) (*model.User, error) {
	user, err := service.UserRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &user, nil
}

func (service *UserService) BlockOrUblock(user *model.User) (*model.User, error) {

	updatedUser, err := service.UserRepo.BlockOrUblock(user)
	if err != nil {
		return nil, fmt.Errorf("couldn't do block/unblock operation")
	}
	return &updatedUser, nil
}

func (service *UserService) FindEmail(user *model.User) string {

	email, err := service.PersonService.FindEmail(user.Id)
	if err != nil {
		return "Invalid email"
	}
	return email
}

func (service *UserService) FindAllUsers() ([]model.User, error) {
	return service.UserRepo.FindAll()
}

func (service *UserService) Registration(registration *model.Registration, token *string, item *bool) error {

	newUser := model.User{
		Username:          registration.Username,
		Password:          registration.Password,
		Role:              model.ParseUserRole(registration.Role),
		VerificationToken: *token,
		IsActive:          *item,
	}
	newPerson := model.Person{
		Name:    registration.Name,
		Surname: registration.Surname,
		Email:   registration.Email,
	}

	err := service.UserRepo.RegisterUser(&newUser)
	if err != nil {
		return err
	}

	newPerson.UserId = newUser.Id

	err = service.UserRepo.RegisterPerson(&newPerson)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetAndVerifyUserByToken(token *string) (*model.User, error) {
	user, err := service.UserRepo.GetUserByToken(token)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with token %s not found", *token))
	}
	*&user.IsActive = true
	updatedUser, err := service.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (service *UserService) UpdateUser(user *model.User) (*model.User, error) {
	updatedUser, err := service.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

/* bez tokena
func (service *UserService) Register(accountRegDto *dto.AccountRegistrationDto) (*dto.AuthenticationTokensDto, error) {
	users, err := service.UserRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("couldn't register")
	}
	if len(users) != 0 {
		for _, user := range users {
			if user.Username == accountRegDto.Username {
				return nil, fmt.Errorf("username already exists")
			}

			email, _ := service.PersonService.FindEmail(user.Id)
			if email == accountRegDto.Email {
				return nil, fmt.Errorf("email already exists")
			}
		}
	}
	service.CreateUserAndPerson(accountRegDto)
	token := dto.AuthenticationTokensDto{
		Id:          8,
		AccessToken: "lala",
	}
	return &token, nil
}

func (service *UserService) CreateUserAndPerson(accountRegDto *dto.AccountRegistrationDto) {
	verificationToken := GenerateUniqueVerificationToken()
	user := model.User{
		Username:          accountRegDto.Username,
		Password:          accountRegDto.Password,
		Role:              model.Tourist,
		IsActive:          true,
		VerificationToken: verificationToken,
	}
	id, _ := service.UserRepo.CreateUser(&user)

	person := model.Person{
		UserId:  id,
		Name:    accountRegDto.Name,
		Surname: accountRegDto.Surname,
		Email:   accountRegDto.Email,
	}
	service.PersonService.Create(&person)
	//generisanje id-a automatski, emailsender, vratiti token
}

func GenerateUniqueVerificationToken() string {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		// Handle error
		return ""
	}

	verificationToken := hex.EncodeToString(tokenBytes)
	return verificationToken
}
*/
