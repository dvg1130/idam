package api

import (
	"database/sql"
	"net/http"
	"tiered-service-backend/internal/middleware"
	"tiered-service-backend/internal/validators"

	"go.uber.org/zap"
)

type Handlers struct {
	Handler  http.HandlerFunc
	Login    http.HandlerFunc
	Register http.HandlerFunc
	Logout   http.HandlerFunc

	DashboardPOST       http.HandlerFunc
	DashboardGET        http.HandlerFunc
	SnakeProfile        http.HandlerFunc
	SnakeUpdate         http.HandlerFunc
	SnakeDelete         http.HandlerFunc
	SnakeFeedAdd        http.HandlerFunc
	SnakeFeedGet        http.HandlerFunc
	SnakeFeedDelete     http.HandlerFunc
	SnakeFeedUpdate     http.HandlerFunc
	SnakeHealthAdd      http.HandlerFunc
	SnakeHealthGet      http.HandlerFunc
	SnakeHealthUpdate   http.HandlerFunc
	SnakeHealthDelete   http.HandlerFunc
	SnakeBreedingAdd    http.HandlerFunc
	SnakeBreedingGetOne http.HandlerFunc
	SnakeBreedingGetAll http.HandlerFunc
	SnakeBreedingUpdate http.HandlerFunc
	SnakeBreedingDelete http.HandlerFunc
	Admin               http.HandlerFunc
	Submit              http.HandlerFunc
	Refresh             http.HandlerFunc
	Health              http.HandlerFunc
	Logger              *zap.Logger
	S                   *sql.DB
}

func InitRoutes(router *http.ServeMux, h *Handlers) {

	//routes
	router.HandleFunc("/", h.Handler)

	router.HandleFunc("/login", validators.Method(http.MethodPost, h.Login))

	router.HandleFunc("/register", validators.Method(http.MethodPost, h.Register))

	router.HandleFunc("/logout", validators.Method(http.MethodPost, h.Logout))

	router.Handle("/dashboard/post",
		middleware.AuthMiddleware(
			middleware.RecordsLimiter(h.S)(
				validators.Method(http.MethodPost, http.HandlerFunc(h.DashboardPOST)),
			),
		),
	)
	//dashboardGet - basic
	router.Handle("/dashboard/get",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.DashboardGET)),
			),
		),
	)
	//snake profile - basic
	router.Handle("/dashboard/snake",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.SnakeProfile)),
			),
		),
	)
	//snake update- basic
	router.Handle("/dashboard/snake/update",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPatch, http.HandlerFunc(h.SnakeUpdate)),
			),
		),
	)

	// delete snake
	router.Handle("/dashboard/snake/delete",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodDelete, http.HandlerFunc(h.SnakeDelete)),
			),
		),
	)

	//feeds

	// get snake feeds by snake
	router.Handle("/dashboard/snake/feed",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.SnakeFeedGet)),
			),
		),
	)

	// post snake feed
	router.Handle("/dashboard/snake/feed/post",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPost, http.HandlerFunc(h.SnakeFeedAdd)),
			),
		),
	)

	// update snake feed
	router.Handle("/dashboard/snake/feed/update",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPatch, http.HandlerFunc(h.SnakeFeedUpdate)),
			),
		),
	)

	// delete snake feed

	router.Handle("/dashboard/snake/feed/delete",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodDelete, http.HandlerFunc(h.SnakeFeedDelete)),
			),
		),
	)

	//health

	// post snake health
	router.Handle("/dashboard/snake/health/post",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPost, http.HandlerFunc(h.SnakeHealthAdd)),
			),
		),
	)

	// get snake health
	router.Handle("/dashboard/snake/health",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.SnakeHealthGet)),
			),
		),
	)

	// update snake health
	router.Handle("/dashboard/snake/health/update",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPatch, http.HandlerFunc(h.SnakeHealthUpdate)),
			),
		),
	)

	//delete snake health record
	router.Handle("/dashboard/snake/health/delete",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodDelete, http.HandlerFunc(h.SnakeHealthDelete)),
			),
		),
	)

	// breeding

	// post breeding event
	router.Handle("/dashboard/breeding/post",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPost, http.HandlerFunc(h.SnakeBreedingAdd)),
			),
		),
	)

	// get all breedings
	router.Handle("/dashboard/breeding/all",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.SnakeBreedingGetAll)),
			),
		),
	)

	// get one breeding event
	router.Handle("/dashboard/breeding/one",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodGet, http.HandlerFunc(h.SnakeBreedingGetOne)),
			),
		),
	)

	// get one breeding event
	router.Handle("/dashboard/breeding/update",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodPatch, http.HandlerFunc(h.SnakeBreedingUpdate)),
			),
		),
	)

	//get delete breeding event
	router.Handle("/dashboard/breeding/delete",
		middleware.AuthMiddleware(
			middleware.RequireRole("basic")(
				validators.Method(http.MethodDelete, http.HandlerFunc(h.SnakeBreedingDelete)),
			),
		),
	)

	//admin
	router.Handle("/admin",
		middleware.AuthMiddleware(
			middleware.RequireRole("admin")(

				validators.Method(http.MethodGet, http.HandlerFunc(h.Admin)),
			),
		),
	)

	// submit
	router.Handle("/submit",
		middleware.LoggingMiddleware(h.Logger)(middleware.AuthMiddleware(
			middleware.PayloadLimiter(
				validators.Method(http.MethodPost, http.HandlerFunc(h.Submit)),
			),
		),
		),
	)

	//refresh token
	router.Handle("/token/refresh", validators.Method(http.MethodPost, http.HandlerFunc(h.Refresh)))

	//health
	router.Handle("/health", middleware.LoggingMiddleware(h.Logger)(http.HandlerFunc(h.Health)))

	//helper

}
