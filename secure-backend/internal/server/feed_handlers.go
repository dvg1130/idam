package server

import "net/http"

// GET snake feed recs
func (s *Server) SnakeFeedGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeFeedGet"))
}

// POST snake feed rec
func (s *Server) SnakeFeedPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeFeedPost"))
}

//UPDATE snake feed rec

func (s *Server) SnakeFeedUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeFeedUpdate"))
}

// DELETE snake feed rec
func (s *Server) SnakeFeedDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeFeedDelete"))
}
