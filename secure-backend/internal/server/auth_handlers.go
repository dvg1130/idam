package server

import (
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/auth"
	"github.com/dvg1130/Portfolio/secure-backend/internal/helpers"
	validator "github.com/dvg1130/Portfolio/secure-backend/internal/validator/auth"
	"github.com/dvg1130/Portfolio/secure-backend/models"
	authdb "github.com/dvg1130/Portfolio/secure-backend/repo/auth_db"
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

	// existing username check
	exists, err := validator.ExistingUser(s.AUTH_DB, req.Username)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}

	//hash pw
	hashedPW, err := auth.HashPW(req.Password)
	if err != nil {
		http.Error(w, "error hashing password", http.StatusInternalServerError)
		return
	}

	//db insert
	_, err = s.AUTH_DB.Exec(authdb.RegisterUser, req.Username, hashedPW)
	if err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	//successfull registration
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("User registered successfully"))
}

// logout
func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to Logout"))
}

// refresh token
func (s *Server) TokenRefresh(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to Logout"))
}
