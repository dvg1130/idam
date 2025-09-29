package api

import (
	"database/sql"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/middleware"
	validator "github.com/dvg1130/Portfolio/secure-backend/internal/validator/util"
	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// init routes auth
func InitRoutes_Data(router *http.ServeMux, S *sql.DB, h *models.DataHandlers) {

	//snake routes
	router.Handle("/dashboard",
		middleware.AuthMiddleware(
			validator.Method(http.MethodGet, http.HandlerFunc(h.SnakeGetAll)),
		),
	)

	router.Handle("/dashboard/snake",
		middleware.AuthMiddleware(
			validator.Method(http.MethodGet, http.HandlerFunc(h.SnakeGetOne)),
		),
	)

	router.Handle("/dashboard/snake/post", middleware.AuthMiddleware(
		middleware.RecordsLimiter(h.S)(
			validator.Method(http.MethodPost, http.HandlerFunc(h.SnakePost)),
		),
	),
	)

	router.Handle("/dashboard/snake/update",
		middleware.AuthMiddleware(
			validator.Method(http.MethodPatch, http.HandlerFunc(h.SnakeUpdate)),
		),
	)

	router.Handle("/dashboard/snake/delete",
		middleware.AuthMiddleware(
			validator.Method(http.MethodDelete, http.HandlerFunc(h.SnakeDelete)),
		),
	)

	//feed routes
	router.Handle("/dashboard/snake/feeds",
		middleware.AuthMiddleware(
			validator.Method(http.MethodGet, http.HandlerFunc(h.SnakeFeedGet)),
		),
	)

	router.Handle("/dashboard/snake/feeds/post",
		middleware.AuthMiddleware(
			validator.Method(http.MethodPost, http.HandlerFunc(h.SnakeFeedPost)),
		),
	)

	router.Handle("/dashboard/snake/feeds/update",
		middleware.AuthMiddleware(
			validator.Method(http.MethodPatch, http.HandlerFunc(h.SnakeFeedUpdate)),
		),
	)

	router.Handle("/dashboard/snake/feeds/delete",
		middleware.AuthMiddleware(
			validator.Method(http.MethodDelete, http.HandlerFunc(h.SnakeFeedDelete)),
		),
	)

	// //health routes
	router.Handle("/dashboard/snake/health", middleware.AuthMiddleware(
		validator.Method(http.MethodGet,
			http.HandlerFunc(h.SnakeHealthGet)),
	),
	)

	router.Handle("/dashboard/snake/health/post", middleware.AuthMiddleware(
		validator.Method(http.MethodPost,
			http.HandlerFunc(h.SnakeHealthPost)),
	),
	)

	router.Handle("/dashboard/snake/health/update", middleware.AuthMiddleware(
		validator.Method(http.MethodPatch,
			http.HandlerFunc(h.SnakeHealthUpdate)),
	),
	)

	router.Handle("/dashboard/snake/health/delete", middleware.AuthMiddleware(
		validator.Method(http.MethodDelete,
			http.HandlerFunc(h.SnakeHealthDelete)),
	),
	)

	//breeding routes
	router.Handle("/dashboard/breeding/all", middleware.AuthMiddleware(
		validator.Method(http.MethodGet,
			http.HandlerFunc(h.SnakeBreedGetAll)),
	),
	)

	router.Handle("/dashboard/breeding/one", middleware.AuthMiddleware(
		validator.Method(http.MethodGet,
			http.HandlerFunc(h.SnakeBreedGetOne)),
	),
	)

	router.Handle("/dashboard/breeding/post", middleware.AuthMiddleware(
		validator.Method(http.MethodPost,
			http.HandlerFunc(h.SnakeBreedPost)),
	),
	)

	router.Handle("/dashboard/breeding/update", middleware.AuthMiddleware(
		validator.Method(http.MethodPatch,
			http.HandlerFunc(h.SnakeBreedUpdate)),
	),
	)

	router.Handle("/dashboard/breeding/delete", middleware.AuthMiddleware(
		validator.Method(http.MethodDelete,
			http.HandlerFunc(h.SnakeBreedDelete)),
	),
	)

}
