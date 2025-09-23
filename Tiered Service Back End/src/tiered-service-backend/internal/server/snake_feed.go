package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tiered-service-backend/internal/auth"
)

// POST /snake/feed/add
func (s *Server) snakeFeedAdd(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	defer r.Body.Close()
	var req struct {
		Sid      string `json:"sid"`       // snake id (user-given)
		FeedDate string `json:"feed_date"` // YYYY-MM-DD
		PreyType string `json:"prey_type"`
		PreySize string `json:"prey_size"`
		Notes    string `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 1️⃣ Find the snake’s suid using owner_uuid + sid
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		userUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2️⃣ Insert feeding record with suid as foreign key
	_, err = s.DATA_DB.Exec(
		`INSERT INTO feeding
            (snake_uuid, sid, feed_date, prey_type, prey_size, notes)
         VALUES (?, ?, ?, ?, ?, ?)`,
		suid, req.Sid, req.FeedDate, req.PreyType, req.PreySize, req.Notes,
	)
	if err != nil {
		http.Error(w, "error inserting feeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"feeding record added"}`))
}

// GET snake feeds by snake
func (s *Server) snakeFeedGet(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	// Expect {"sid":"<snake-id>"} in request body
	defer r.Body.Close()
	var req struct {
		Sid string `json:"sid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// First, get suid for this user’s snake
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		userUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error retrieving snake id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Now get all feeding records for that snake
	rows, err := s.DATA_DB.Query(
		`SELECT sid, feed_date, prey_type, prey_size, notes
           FROM feeding
          WHERE snake_uuid = ?
          ORDER BY feed_date DESC`,
		suid,
	)
	if err != nil {
		http.Error(w, "error retrieving feeding records: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Feeding struct {
		Sid      string `json:"sid"`
		FeedDate string `json:"feed_date"`
		PreyType string `json:"prey_type"`
		PreySize string `json:"prey_size"`
		Notes    string `json:"notes"`
	}

	var feeds []Feeding
	for rows.Next() {
		var f Feeding
		if err := rows.Scan(&f.Sid,
			&f.FeedDate, &f.PreyType, &f.PreySize, &f.Notes); err != nil {
			http.Error(w, "error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		feeds = append(feeds, f)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "row iteration error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(feeds); err != nil {
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
	}
}

// Update feed record by sid uuid and day
// PATCH /snake/feed/update
func (s *Server) snakeFeedUpdate(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	defer r.Body.Close()
	var req struct {
		Sid      string  `json:"sid"`
		FeedDate string  `json:"feed_date"` // YYYY-MM-DD
		PreyType *string `json:"prey_type,omitempty"`
		PreySize *string `json:"prey_size,omitempty"`
		Notes    *string `json:"notes,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate: need sid + feed_date + at least one update field
	if req.Sid == "" || req.FeedDate == "" {
		http.Error(w, "sid and feed_date are required", http.StatusBadRequest)
		return
	}
	if req.PreyType == nil && req.PreySize == nil && req.Notes == nil {
		http.Error(w, "at least one field to update must be provided", http.StatusBadRequest)
		return
	}

	// 1️⃣ Get suid from snakes
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		userUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error retrieving snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2️⃣ Build dynamic update query
	updates := []string{}
	args := []interface{}{}

	if req.PreyType != nil {
		updates = append(updates, "prey_type = ?")
		args = append(args, *req.PreyType)
	}
	if req.PreySize != nil {
		updates = append(updates, "prey_size = ?")
		args = append(args, *req.PreySize)
	}
	if req.Notes != nil {
		updates = append(updates, "notes = ?")
		args = append(args, *req.Notes)
	}

	// Add WHERE args (suid + feed_date)
	args = append(args, suid, req.FeedDate)

	query := fmt.Sprintf(
		"UPDATE feeding SET %s WHERE snake_uuid = ? AND feed_date = ?",
		strings.Join(updates, ", "),
	)

	// 3️⃣ Execute update
	result, err := s.DATA_DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "error updating feeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "could not confirm update: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "feeding record not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"feeding record updated"}`))
}

// delete feed record

type SnakeFeedDeleteRequest struct {
	Sid  string `json:"sid"`       // snake ID or name
	Date string `json:"feed_date"` // feeding record date (match the format stored in DB)
}

func (s *Server) snakeFeedDelete(w http.ResponseWriter, r *http.Request) {
	// ---- 1. Verify JWT claims ----
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	// ---- 2. Parse and validate request body ----
	var req SnakeFeedDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Sid == "" || req.Date == "" {
		http.Error(w, "sid and date are required", http.StatusBadRequest)
		return
	}

	// ---- 3. Get suid for this snake ----
	var suid string
	err := s.DATA_DB.QueryRow(
		"SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?",
		userUUID, req.Sid,
	).Scan(&suid)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "snake not found", http.StatusNotFound)
			return
		}
		http.Error(w, "error retrieving snake", http.StatusInternalServerError)
		return

	}

	// ---- 4. Delete feeding record ----
	res, err := s.DATA_DB.Exec(
		"DELETE FROM feeding WHERE snake_uuid = ? AND feed_date = ?",
		suid, req.Date,
	)

	if err != nil {
		http.Error(w, "error deleting feeding record", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "feeding record not found", http.StatusNotFound)
		return
	}

	// ---- 5. Respond success ----
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"feeding record deleted successfully"}`))
}
