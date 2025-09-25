package server

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/helpers"
	"github.com/dvg1130/Portfolio/secure-backend/internal/validator"
	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// entry
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to server"))

}

// login
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to Login"))

}

// register
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {

	//decode body
	req, err := helpers.DecodeBody[models.Credentials](w, r)
	if err != nil {
		return
	}

	//len check
	if err := validator.AuthLenCheck(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

}

// logout
func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to Logout"))
}

// refresh token
func (s *Server) TokenRefresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to Logout"))
}
