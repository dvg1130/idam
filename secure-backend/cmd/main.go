package main

import (
	"fmt"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/app"
)

// init server
func main() {

	server := app.AppServer()
	err := http.ListenAndServe(":8003", server.Router)
	if err != nil {
		fmt.Println("error starting server")
	}
}
