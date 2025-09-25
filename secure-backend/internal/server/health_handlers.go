package server

import "net/http"

// GET snake fHealth recs
func (s *Server) SnakeHealthGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeHealthGet"))
}

// POST snake Health rec
func (s *Server) SnakeHealthPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeHealthPost"))
}

//UPDATE snake Health rec

func (s *Server) SnakeHealthUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeHealthUpdate"))
}

// DELETE snake Health rec
func (s *Server) SnakeHealthDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeHealthDelete"))
}
