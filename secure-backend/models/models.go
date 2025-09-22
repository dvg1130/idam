package models

import "net/http"

type Server struct {
	Router *http.ServeMux
	// AUTH_DB *sql.DB
	// Data_DB *sql.DB
	// Redis   *redis.Client
	// Logger  *zap.Logger
}
