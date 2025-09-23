package api

import (
	"net/http"
)

type Handlers struct {
	Handler http.HandlerFunc
	Login   http.HandlerFunc
}

// init routes
func InitRoutes(router *http.ServeMux, h *Handlers) {

	//routes
	router.HandleFunc("/", h.Handler)
	router.HandleFunc("/login", h.Login)
}
