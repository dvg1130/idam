package api

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

func InitRoutes_Admin(router *http.ServeMux, h *models.AdminHandlers) {

	//routes

	router.HandleFunc("/admin/uses/all", h.AdminGetAll)
	router.HandleFunc("/admin/user/one", h.AdminGetOne)
	router.HandleFunc("/admin/user/update", h.AdminUpdate)
}
