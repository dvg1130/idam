package server

import "net/http"

// GET all SnakeBreeding events
func (s *Server) SnakeBreedGetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedGetAll"))
}

// GET one SnakeBreed event
func (s *Server) SnakeBreedGetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedGetOne"))
}

// POST SnakeBreed event
func (s *Server) SnakeBreedPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedPost"))
}

// UPDATE SnakeBreedevent
func (s *Server) SnakeBreedUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedUpdate"))
}

// DELETE SnakeBreed event
func (s *Server) SnakeBreedDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedDelete"))
}
