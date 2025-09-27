package server

import (
	"encoding/json"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/auth"
	"github.com/dvg1130/Portfolio/secure-backend/logs"
	"github.com/dvg1130/Portfolio/secure-backend/models"
)

// GET one user lsit
func (s *Server) AdminGetAll(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	username := claims["username"].(string)
	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		logs.LogEvent(s.Logger, "warn", "Unauthorized access", r, map[string]interface{}{
			"username": username,
			"role":     role,
		})
		return
	}
	//query all users and roles
	rows, err := s.AUTH_DB.Query(`SELECT uuid, username, role FROM users`)
	if err != nil {
		http.Error(w, "database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UUID, &u.Username, &u.Role); err != nil {
			http.Error(w, "row scan failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "iteration error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "encode failed: "+err.Error(), http.StatusInternalServerError)
	}
}

// GET all user lsit
func (s *Server) AdminGetOne(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to AdminGetOne"))
}

// UPDATE user role
func (s *Server) AdminUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to AdminUpdate"))
}
