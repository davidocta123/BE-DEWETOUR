package database

import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Country{}, &models.Trip{}, &models.Transaction{})

	if err != nil {
		panic("Migration Failed")
	}
}
