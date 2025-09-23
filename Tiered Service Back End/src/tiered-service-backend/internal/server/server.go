package server

import (
	"database/sql"
	"net/http"
	"tiered-service-backend/internal/api"
	"tiered-service-backend/internal/middleware"
	"tiered-service-backend/repository/redisdb"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Server struct {
	Router  *http.ServeMux
	DB      *sql.DB
	DATA_DB *sql.DB
	Redis   *redis.Client
	Logger  *zap.Logger
}

func NewServer(db *sql.DB, data_db *sql.DB, logger *zap.Logger) *Server {
	rdb := redisdb.NewClient()

	s := &Server{
		Router: http.NewServeMux(),
		//db
		DB:      db,
		DATA_DB: data_db,
		Redis:   rdb,
		Logger:  logger,
	}
	api.InitRoutes(s.Router, &api.Handlers{
		Handler:  s.handler,
		Login:    s.login,
		Register: s.register,
		Logout:   s.logout,

		DashboardPOST:       s.dashboardPost,
		DashboardGET:        s.dashboardGet,
		SnakeProfile:        s.snakeProfile,
		SnakeUpdate:         s.snakeUpdate,
		SnakeDelete:         s.snakeDelete,
		SnakeFeedAdd:        s.snakeFeedAdd,
		SnakeFeedGet:        s.snakeFeedGet,
		SnakeFeedDelete:     s.snakeFeedDelete,
		SnakeFeedUpdate:     s.snakeFeedUpdate,
		SnakeHealthAdd:      s.snakeHealthAdd,
		SnakeHealthGet:      s.snakeHealthGet,
		SnakeHealthUpdate:   s.snakeHealthUpdate,
		SnakeHealthDelete:   s.snakeHealthDelete,
		SnakeBreedingAdd:    s.snakeBreedAdd,
		SnakeBreedingGetOne: s.snakeBreedingGetOne,
		SnakeBreedingGetAll: s.snakeBreedGetAll,
		SnakeBreedingUpdate: s.snakeBreedUpdate,
		SnakeBreedingDelete: s.snakeBreedingDelete,

		Admin:   s.admin,
		Submit:  s.submit,
		Refresh: s.refresh,
		Health:  s.health,
		Logger:  s.Logger,
		S:       s.DATA_DB,
	})

	s.Router = ServeMuxWrapper(
		s.Router,
		middleware.SecurityHeaders,
		middleware.LoggingMiddleware(logger),
	)

	return s
}
