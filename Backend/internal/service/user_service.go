package service

import (
	"forge/internal/models"
	"forge/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (u UserService) RegisterUser(userStruct models.User) error {

	return u.userRepository.CreateUser(&userStruct)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}
