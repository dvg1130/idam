package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// config struct
type AuthConfigStruct struct {
	DATABASE_URL   string
	DB_DRIVER      string
	JWT_SECRET_KEY string
	PORT           string
	REDIS_ADDR     string
}

type DataConfigStruct struct {
	DATABASE_URL string
	DB_DRIVER    string
	PORT         string
}

// init & load vars
var AuthConfig AuthConfigStruct
var DataConfig DataConfigStruct

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("error loading env file")
	}

	AuthConfig = AuthConfigStruct{
		DATABASE_URL:   os.Getenv("DATABASE_URL"),
		DB_DRIVER:      os.Getenv("DB_DRIVER"),
		PORT:           os.Getenv("PORT"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
		REDIS_ADDR:     os.Getenv("REDIS_ADDR"),
	}

	DataConfig = DataConfigStruct{
		DATABASE_URL: os.Getenv("Data_DATABASE_URL"),
		DB_DRIVER:    os.Getenv("Data_DB_DRIVER"),
		PORT:         os.Getenv("Data_PORT"),
	}

}
