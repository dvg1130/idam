package server

import "net/http"

// GET all snakes
func (s *Server) SnakeGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeGetAll"))
}

// GET one snake
func (s *Server) SnakeGetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeGetOne"))
}

// POST snake
func (s *Server) SnakePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakePost"))
}

// UPDATE snake
func (s *Server) SnakeUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeUpdate"))
}

// DELETE snake
func (s *Server) SnakeDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeDelete"))
}
