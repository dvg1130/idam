package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tiered-service-backend/internal/auth"
)

type Snake struct {
	SnakeId string `json:"sid"`
	Species string `json:"species"`
	Sex     string `json:"sex"`
	Age     int    `json:"age"`
	Genes   string `json:"genes"`
	Notes   string `json:"notes"`
}

type Snakes struct {
	Sid     string `json:"sid"`
	Species string `json:"species"`
}

// GET dashboard - basic level user
func (s *Server) dashboardGet(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}

	uuid := claims["uuid"].(string)

	// Query all snakes for this owner
	rows, err := s.DATA_DB.Query(
		"SELECT sid, species FROM snakes WHERE owner_uuid = ?", uuid,
	)
	if err != nil {
		http.Error(w, "error retrieving snakes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var snakeList []Snakes
	for rows.Next() {
		var sItem Snakes
		if err := rows.Scan(&sItem.Sid, &sItem.Species); err != nil {
			http.Error(w, "error scanning row", http.StatusInternalServerError)
			return
		}
		snakeList = append(snakeList, sItem)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "error iterating rows", http.StatusInternalServerError)
		return
	}

	// Return JSON list
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(snakeList); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

// POST dashboard free tier
func (s *Server) dashboardPost(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return

	}
	uuid := claims["uuid"].(string)
	role := claims["role"].(string)

	//decode json
	var snake Snake
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&snake); err != nil {
		http.Error(w, "invalid request body"+err.Error(), http.StatusBadRequest)
		return
	}

	// existing snake check (unique sid per user)
	var exists bool
	err := s.DATA_DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM snakes WHERE sid = ? AND owner_uuid = ?)",
		snake.SnakeId, uuid,
	).Scan(&exists)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Snake ID already exists for this user", http.StatusConflict)
		return
	}

	//rbac middleware will wrap post method at router level, additional role vs max allowed check.. if then error, else then return

	// insert into DB (using prepared statement)
	_, err = s.DATA_DB.Exec(
		//maybe add snakeuuid for unique snake id across all users to prevent data leak
		"INSERT INTO snakes (owner_uuid, sid, species, sex, age, genes, notes) VALUES (?, ?, ?, ?, ?, ?, ?)",
		uuid,
		snake.SnakeId,
		snake.Species,
		snake.Sex,
		snake.Age,
		snake.Genes,
		snake.Notes,
	)

	if err != nil {
		http.Error(w, "error creating new snake", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Upload Successful", uuid, role)

}
