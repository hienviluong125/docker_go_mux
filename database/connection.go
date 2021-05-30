package database

import (
	"fmt"
	"os"

	"github.com/hienviluong125/docker_go_mux/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	connection, err := gorm.Open(postgres.Open(dbs), &gorm.Config{})

	if err != nil {
		panic("Database connection error")
	}

	DB = connection

	connection.AutoMigrate(&models.Post{})
}
