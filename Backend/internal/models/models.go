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
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Instructor struct {
	UserID         string `json:"user_id" gorm:"index"`                              // 'UserID' es la clave foránea
	User           User   `json:"user" gorm:"foreignKey:UserID;references:Username"` // Relación con User
	Specialization string `json:"specialization"`
}

type Trainee struct {
	UserID string  `json:"user_id" gorm:"index"`                              // 'UserID' es la clave foránea
	User   User    `json:"user" gorm:"foreignKey:UserID;references:Username"` // Relación con User
	Weight float32 `json:"weight"`
	Height int16   `json:"height"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
