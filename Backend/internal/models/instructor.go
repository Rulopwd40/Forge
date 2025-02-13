package models

type Instructor struct {
	UserID         string `json:"user_id" gorm:"index"`                              // 'UserID' es la clave foránea
	User           User   `json:"user" gorm:"foreignKey:UserID;references:Username"` // Relación con User
	Specialization string `json:"specialization"`
}
