package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Username  string `json:"username" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"password"`
	Level     int    `json:"level"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
