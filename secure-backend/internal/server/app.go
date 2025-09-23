package server

import (
	"database/sql"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/api"
)

type Server struct {
	Router  *http.ServeMux
	AUTH_DB *sql.DB
	Data_DB *sql.DB
	// Redis   *redis.Client
	// Logger  *zap.Logger
}

func AppServer(auth_db *sql.DB, data_db *sql.DB) *Server {

	s := &Server{
		Router:  http.NewServeMux(),
		AUTH_DB: auth_db,
		Data_DB: data_db,
	}

	api.InitRoutes(s.Router, &api.Handlers{
		Handler: s.Handler,
		Login:   s.Login,
	})

	return s
}
