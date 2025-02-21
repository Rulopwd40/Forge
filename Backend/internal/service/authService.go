package service

import (
	"errors"
	"forge/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IAuthService interface {
	Login(loginData models.LoginRequest) (string, error)
}

type AuthService struct {
	UserService     IUserService
	PasswordService IPasswordService
}

func NewAuthService(userService IUserService, passwordService IPasswordService) IAuthService {
	return &AuthService{userService, passwordService}
}

func (a AuthService) Login(loginData models.LoginRequest) (string, error) {

	user, err := a.UserService.GetUser(loginData.Username, loginData.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare password
	err = a.PasswordService.CompareHashAndPassword(user.Password, loginData.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("error signing token")
	}

	return tokenString, nil
}
