package service

import (
	"errors"
	"forge/internal/models"
	"forge/internal/repository"
)

type IUserService interface {
	RegisterUser(userStruct models.User) error
	GetUser(username string, email string) (models.User, any)
	GetUserData(username string) (models.User, error)
}

type UserService struct {
	UserRepository  repository.IUserRepository
	PasswordService IPasswordService
}

func NewUserService(userRepository repository.IUserRepository, passwordService IPasswordService) IUserService {
	return UserService{userRepository, passwordService}
}

func (u UserService) RegisterUser(userStruct models.User) error {
	existingUser := u.UserRepository.CheckUser(&userStruct)
	if existingUser != nil {
		return errors.New("record not found")
	}

	password, err := u.PasswordService.GenerateHashPassword(userStruct.Password)
	if err != nil {
		return err
	}

	userStruct.Password = password

	return u.UserRepository.CreateUser(&userStruct)
}

func (u UserService) GetUser(username string, email string) (models.User, any) {
	if username != "" {
		return u.UserRepository.GetUserByUsername(username)
	}
	if email != "" {
		return u.UserRepository.GetUserByEmail(email)
	}
	return models.User{}, errors.New("no username or email provided")
}

func (u UserService) GetUserData(username string) (models.User, error) {
	if username != "" {
		user, err := u.UserRepository.GetUserByUsername(username)

		if err != nil {
			return models.User{}, errors.New("record not found")
		}

		return models.User{
			Username: user.Username,
			Name:     user.Name,
			Level:    user.Level,
			Email:    user.Email,
		}, nil
	}
	return models.User{}, errors.New("no username provided")
}
