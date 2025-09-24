package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfigStruct struct {
	PORT   string
	HOST   string
	GOMODE string
}

func ServerConfig() ServerConfigStruct {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.Server = ServerConfigStruct{
		PORT:   os.Getenv("PORT"),
		HOST:   os.Getenv("HOST"),
		GOMODE: os.Getenv("GOMODE"),
	}

	return AppConfig.Server
}
