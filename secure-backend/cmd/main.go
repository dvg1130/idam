package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/app"
	authdb "github.com/dvg1130/Portfolio/secure-backend/repo/auth_db"
)

// init server
func main() {

	auth_db, err := authdb.AuthDBClient()
	if err != nil {
		log.Fatal("failed to connect to auth db", err)

		defer auth_db.Close()
	}

	server := app.AppServer(auth_db)
	http.ListenAndServe(":8003", server.Router)
	if err != nil {
		fmt.Println("error starting server")
	}
}
