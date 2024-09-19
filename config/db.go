package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connectdb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")
	databaseport := os.Getenv("DATABASEPORT")
	database := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", password, database, databaseport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to mysql database")
	} else {
		fmt.Println("Connected to db")
	}

	return db
}

func Disconnectdb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection to database")
	}

	dbSQL.Close()
}
