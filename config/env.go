package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var EnvModule = fx.Provide(LoadEnv)

type EnvVars struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func LoadEnv() EnvVars {
	godotenv.Load()
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return EnvVars{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DbName:   dbname,
	}
}
