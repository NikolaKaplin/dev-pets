package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfigStruct struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
}

func DBConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.DB = DBConfigStruct{
		DBHost:         os.Getenv("DB_HOST"),
		DBUserName:     os.Getenv("DB_USERNAME"), // Исправлено: DB_USERNAME
		DBUserPassword: os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBPort:         os.Getenv("DB_PORT"),
	}
}
