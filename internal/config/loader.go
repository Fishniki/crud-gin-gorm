package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DataBase struct {
	Database string
}

func Get() *DataBase {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file in Config")
	}

	return &DataBase{
		Database: os.Getenv("DB_URL"),
	}
}