package api

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// init routes
func InitRouter(router *http.ServeMux, h *models.Handlers) {

	//routes
	router.HandleFunc("/", h.Handler)
	router.HandleFunc("/login", h.Login)
}
