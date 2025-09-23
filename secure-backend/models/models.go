package models

import (
	"database/sql"
	"net/http"
)

type Server struct {
	Router  *http.ServeMux
	AUTH_DB *sql.DB
	// Data_DB *sql.DB
	// Redis   *redis.Client
	// Logger  *zap.Logger
}

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
