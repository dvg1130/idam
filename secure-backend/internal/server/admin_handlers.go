package server

import "net/http"

// GET one user lsit
func (s *Server) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to AdminGetAll"))
}

// GET all user lsit
func (s *Server) AdminGetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to AdminGetOne"))
}

// UPDATE user role
func (s *Server) AdminUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to AdminUpdate"))
}
