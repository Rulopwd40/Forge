package database

import (
	"log"

	"forge/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/forge?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set the directory for migrations
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.Trainee{}, &models.Instructor{})

	return db
}
