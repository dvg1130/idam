package api

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// init routes auth
func InitRoutes_Data(router *http.ServeMux, h *models.DataHandlers) {

	//snake routes
	router.HandleFunc("/dashboard", h.SnakeGetAll)
	router.HandleFunc("/dashboard/snake", h.SnakeGetOne)
	router.HandleFunc("/dashboard/snake/post", h.SnakePost)
	router.HandleFunc("/dashboard/snake/update", h.SnakeUpdate)
	router.HandleFunc("/dashboard/snake/delete", h.SnakeDelete)

	//feed routes
	router.HandleFunc("/dashboard/snake/feeds", h.SnakeFeedGet)
	router.HandleFunc("/dashboard/snake/feeds/post", h.SnakeFeedPost)
	router.HandleFunc("/dashboard/snake/feeds/update", h.SnakeFeedUpdate)
	router.HandleFunc("/dashboard/snake/feeds/delete", h.SnakeFeedDelete)

	// //health routes
	router.HandleFunc("/dashboard/snake/health", h.SnakeHealthGet)
	router.HandleFunc("/dashboard/snake/health/post", h.SnakeHealthPost)
	router.HandleFunc("/dashboard/snake/health/update", h.SnakeHealthUpdate)
	router.HandleFunc("/dashboard/snake/health/delete", h.SnakeHealthDelete)

	//breeding routes
	router.HandleFunc("/dashboard/breeding/all", h.SnakeBreedGetAll)
	router.HandleFunc("/dashboard/breeding/one", h.SnakeBreedGetOne)
	router.HandleFunc("/dashboard/breeding/post", h.SnakeBreedPost)
	router.HandleFunc("/dashboard/breeding/update", h.SnakeBreedUpdate)
	router.HandleFunc("/dashboard/breeding/delete", h.SnakeBreedDelete)

}
