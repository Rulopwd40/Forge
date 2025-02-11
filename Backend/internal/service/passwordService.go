package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type IPasswordService interface {
	GenerateHashPassword(password string) (string, error)
	CompareHashAndPassword(userPassword string, loginPassword string) error
}

type PasswordService struct {
}

func NewPasswordService() IPasswordService {
	return &PasswordService{}
}

func (p PasswordService) GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (p PasswordService) CompareHashAndPassword(userPassword string, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(loginPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
