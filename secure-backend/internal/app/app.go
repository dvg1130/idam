package app

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

func AppServer() *models.Server {
	s := &models.Server{
		Router: http.NewServeMux(),
	}

	return s
}
