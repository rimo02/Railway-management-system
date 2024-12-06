package database

import (
	"Railway-management-system/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:rimo398@tcp(127.0.0.1:3306)/irctc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Train{}, &models.User{}, &models.Booking{})
	return db
}
