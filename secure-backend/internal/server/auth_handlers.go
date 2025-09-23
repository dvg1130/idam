package server

import (
	"net/http"
)

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successful connection to server"))

}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successful connection to Login"))

}
