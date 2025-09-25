package api

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// type AuthHandlers struct {
// 	//auth
// 	Handler  http.HandlerFunc
// 	Login    http.HandlerFunc
// 	Register http.HandlerFunc
// 	Logout   http.HandlerFunc
// }

// init routes auth
func InitRoutes_Auth(router *http.ServeMux, h *models.AuthHandlers) {

	//routes
	router.HandleFunc("/", h.Handler)
	router.HandleFunc("/login", h.Login)
	router.HandleFunc("/register", h.Register)
	router.HandleFunc("/logout", h.Logout)
}
