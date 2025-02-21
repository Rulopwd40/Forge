package repository

import (
	"forge/internal/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUsers() ([]models.User, error)
	CreateUser(user *models.User) error
	CheckUser(user *models.User) error
	GetUserByEmail(email string) (models.User, any)
	GetUserByUsername(username string) (models.User, any)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// Create

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// Get
func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, any) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, any) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) CheckUser(user *models.User) error {
	err := r.db.Where("username = ? OR email = ?", user.Username, user.Email).First(&models.User{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}
