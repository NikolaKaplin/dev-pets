package config

import (
	"golang.org/x/oauth2"
)

type Config struct {
	GithubLoginConfig oauth2.Config
	GoogleLoginConfig oauth2.Config
	Server            ServerConfigStruct
	DB                DBConfigStruct
}

var AppConfig Config

func Load() {
	DBConfig()
	GithubConfig()
	GoogleConfig()
	ServerConfig()
}
