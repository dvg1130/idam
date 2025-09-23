package app

import (
	"database/sql"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

func AppServer(auth_db *sql.DB) *models.Server {
	s := &models.Server{
		Router:  http.NewServeMux(),
		AUTH_DB: auth_db,
	}

	return s
}
