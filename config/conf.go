package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_HOST    string
	API_PORT    string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var ENV Config

func Init_config() {
	godotenv.Load()

	ENV.API_HOST = os.Getenv("API_HOST")
	ENV.API_PORT = os.Getenv("API_PORT")

	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_PORT = os.Getenv("DB_PORT")
	ENV.DB_USER = os.Getenv("DB_USER")
	ENV.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	ENV.DB_NAME = os.Getenv("DB_NAME")
}
