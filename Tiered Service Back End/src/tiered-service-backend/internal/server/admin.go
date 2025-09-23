package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tiered-service-backend/internal/auth"
)

// admin get users func
// adminGet returns a list of all users with their role.
func (s *Server) adminGet(w http.ResponseWriter, r *http.Request) {
	// ✅ Require admin role from the JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	// Query all users and their roles
	rows, err := s.DB.Query(`SELECT uuid, username, role FROM users`)
	if err != nil {
		http.Error(w, "database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type User struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	var users []User
	for rows.Next() {
		var u User
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

//admin update user

func (s *Server) adminUpdate(w http.ResponseWriter, r *http.Request) {
	// ✅ Require admin role from the JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	defer r.Body.Close()
	var req struct {
		Username string `json:"username"`
		OldRole  string `json:"old_role"`
		NewRole  string `json:"new_role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.OldRole == "" || req.NewRole == "" {
		http.Error(w, "username, old_role, and new_role are required", http.StatusBadRequest)
		return
	}

	// 1️⃣ Look up the user’s uuid, ensuring the username/old_role match
	var uuid string
	err := s.DB.QueryRow(
		`SELECT uuid FROM users WHERE username = ? AND role = ?`,
		req.Username, req.OldRole,
	).Scan(&uuid)
	if err == sql.ErrNoRows {
		http.Error(w, "user not found or old role mismatch", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "database query failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2️⃣ Update the role by uuid
	_, err = s.DB.Exec(
		`UPDATE users SET role = ? WHERE uuid = ?`,
		req.NewRole, uuid,
	)
	if err != nil {
		http.Error(w, "failed to update role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"user role updated successfully"}`))
}

//delete user
