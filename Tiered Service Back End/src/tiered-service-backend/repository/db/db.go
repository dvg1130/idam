package db

import (
	"database/sql"
	"fmt"
	"tiered-service-backend/config"

	_ "github.com/go-sql-driver/mysql"
)

// db
func DBClient() (*sql.DB, error) {
	db, err := sql.Open(config.AuthConfig.DB_DRIVER, config.AuthConfig.DATABASE_URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to database")
	return db, nil
}

func DataDBClient() (*sql.DB, error) {
	db, err := sql.Open(config.DataConfig.DB_DRIVER, config.DataConfig.DATABASE_URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to database")
	return db, nil
}
