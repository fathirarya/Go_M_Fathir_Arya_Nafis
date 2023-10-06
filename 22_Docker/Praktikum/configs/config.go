package configs

import (
	"Praktikum/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func InitDB() {
	config := Config{
		DB_Username: "fathir",
		DB_Password: "ayam12345",
		DB_Port:     "3306",
		DB_Host:     "db4free.net",
		DB_Name:     "middleware",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(models.User{}, models.Book{})
}
