package app

import (
	"database/sql"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

func AppServer(auth_db *sql.DB, data_db *sql.DB) *models.Server {
	s := &models.Server{
		Router:  http.NewServeMux(),
		AUTH_DB: auth_db,
		Data_DB: data_db,
	}

	return s
}
