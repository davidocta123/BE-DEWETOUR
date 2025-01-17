package mysql

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connetion ke database
func DatabaseConnection() {
	var err error

	// dsn := "root:@tcp(localhost:3306)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"

	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to database")

}
