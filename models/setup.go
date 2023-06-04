package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB //todo: make not global

func ConnectToDb() {
	err := godotenv.Load(".env") //todo: remove in docker
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_NAME"),
		os.Getenv("PG_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Database connection error:", err)
	} else {
		fmt.Println("Connected to the database")
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("failed to migrate db: ", err)
	}
}
