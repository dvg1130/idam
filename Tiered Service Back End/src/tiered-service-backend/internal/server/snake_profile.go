package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tiered-service-backend/internal/auth"
)

// GET snake profile /snake
func (s *Server) snakeProfile(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	uuid := claims["uuid"].(string)

	defer r.Body.Close()
	var req Snake
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// query a single snake
	var snake Snake
	err := s.DATA_DB.QueryRow(
		`SELECT sid, species, sex, age, genes, notes
         FROM snakes
         WHERE sid = ? AND owner_uuid = ?`,
		req.SnakeId, uuid,
	).Scan(
		&snake.SnakeId, &snake.Species, &snake.Sex,
		&snake.Age, &snake.Genes, &snake.Notes,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error retrieving snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(snake); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}

// delete snake /snake/delete
func (s *Server) snakeDelete(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	uuid := claims["uuid"].(string)

	defer r.Body.Close()
	var req struct {
		SnakeId string `json:"sid"` // expecting {"sid": "..."}
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.DATA_DB.Exec(
		`DELETE FROM snakes
         WHERE sid = ? AND owner_uuid = ?`,
		req.SnakeId, uuid,
	)
	if err != nil {
		http.Error(w, "error deleting snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "could not confirm deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "snake not found or not owned by user", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 – success, no response body
	w.Write([]byte("snakes deleted"))
}

// update snake /snake/update
type UpdateSnakeRequest struct {
	Sid     string  `json:"sid"` // required
	Species *string `json:"species,omitempty"`
	Sex     *string `json:"sex,omitempty"`
	Age     *int    `json:"age,omitempty"`
	Genes   *string `json:"genes,omitempty"`
	Notes   *string `json:"notes,omitempty"`
}

func (s *Server) snakeUpdate(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	uuid := claims["uuid"].(string)

	defer r.Body.Close()

	// decode json into a struct with pointer f

	var req UpdateSnakeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" {
		http.Error(w, "sid is required", http.StatusBadRequest)
		return
	}

	// build the SET clause dynamically
	sets := []string{}
	args := []any{}

	if req.Species != nil {
		sets = append(sets, "species = ?")
		args = append(args, *req.Species)
	}
	if req.Sex != nil {
		sets = append(sets, "sex = ?")
		args = append(args, *req.Sex)
	}
	if req.Age != nil {
		sets = append(sets, "age = ?")
		args = append(args, *req.Age)
	}
	if req.Genes != nil {
		sets = append(sets, "genes = ?")
		args = append(args, *req.Genes)
	}
	if req.Notes != nil {
		sets = append(sets, "notes = ?")
		args = append(args, *req.Notes)
	}

	if len(sets) == 0 {
		http.Error(w, "no fields to update", http.StatusBadRequest)
		return
	}

	// add WHERE arguments: sid and owner_uuid
	args = append(args, req.Sid, uuid)

	query := fmt.Sprintf("UPDATE snakes SET %s WHERE sid = ? AND owner_uuid = ?",
		strings.Join(sets, ", "))

	if _, err := s.DATA_DB.Exec(query, args...); err != nil {
		http.Error(w, "update failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"updated"}`))
}
