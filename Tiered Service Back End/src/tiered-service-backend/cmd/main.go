package main

import (
	"fmt"
	"log"
	"net/http"
	"tiered-service-backend/internal/server"
	"tiered-service-backend/repository"
	"tiered-service-backend/repository/db"
)

// init server
func main() {
	logger := repository.NewLogger()
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Printf("failed to flush logger: %v\n", err)
		}
	}()

	auth_database, err := db.DBClient()
	if err != nil {
		log.Fatal("failed to connectto database", err)

		defer auth_database.Close()

	}

	data_database, err := db.DataDBClient()
	if err != nil {
		log.Fatal("failed to connect to database", err)

		defer data_database.Close()

	}
	server := server.NewServer(auth_database, data_database, logger)
	http.ListenAndServe(":8002", server.Router)
	if err != nil {
		fmt.Println("Error starting server")
	}

}
