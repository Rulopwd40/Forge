package models

type Trainee struct {
	UserID string  `json:"user_id" gorm:"index"`                              // 'UserID' es la clave foránea
	User   User    `json:"user" gorm:"foreignKey:UserID;references:Username"` // Relación con User
	Weight float32 `json:"weight"`
	Height int16   `json:"height"`
}
