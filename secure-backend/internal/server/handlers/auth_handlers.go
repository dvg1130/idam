package handlers

import (
	"fmt"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/models"
)

func handler(s *models.Server, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successful connection to server"))
	fmt.Print()
}
