package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tiered-service-backend/internal/auth"
)

//retrieve health records all

func (s *Server) snakeHealthGet(w http.ResponseWriter, r *http.Request) {
	// 1️⃣ Verify JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	// 2️⃣ Parse request body
	defer r.Body.Close()
	var req struct {
		Sid string `json:"sid"` // user-provided snake id
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" {
		http.Error(w, "sid is required", http.StatusBadRequest)
		return
	}

	// 3️⃣ Find suid for this snake
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

	// 4️⃣ Query all health records for this suid
	rows, err := s.DATA_DB.Query(
		`SELECT check_date, weight, length, topic, notes
		   FROM health
		  WHERE suid = ?
		  ORDER BY check_date DESC`,
		suid,
	)
	if err != nil {
		http.Error(w, "error retrieving health records: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 5️⃣ Collect results
	type HealthRecord struct {
		CheckDate string `json:"check_date"`
		Weight    string `json:"weight"`
		Length    string `json:"length"`
		Topic     string `json:"topic"`
		Notes     string `json:"notes"`
	}
	var records []HealthRecord

	for rows.Next() {
		var hr HealthRecord
		if err := rows.Scan(&hr.CheckDate, &hr.Weight, &hr.Length, &hr.Topic, &hr.Notes); err != nil {
			http.Error(w, "error scanning row: "+err.Error(), http.StatusInternalServerError)
			return
		}
		records = append(records, hr)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "error iterating rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 6️⃣ Return JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(records); err != nil {
		http.Error(w, "failed to encode JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// post health record
func (s *Server) snakeHealthAdd(w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	defer r.Body.Close()
	var req struct {
		Sid       string `json:"sid"`        // snake id (user-given)
		CheckDate string `json:"check_date"` // YYYY-MM-DD, required
		Weight    string `json:"weight"`     // optional
		Length    string `json:"length"`     // optional
		Topic     string `json:"topic"`      // optional
		Notes     string `json:"notes"`      // optional
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// ✅ Require check_date
	if req.CheckDate == "" {
		http.Error(w, "check_date is required", http.StatusBadRequest)
		return
	}

	// ✅ Require at least one optional field
	if req.Weight == "" && req.Length == "" && req.Topic == "" && req.Notes == "" {
		http.Error(w, "at least one of weight, length, topic, or notes must be provided", http.StatusBadRequest)
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

	// 2️⃣ Insert health record
	_, err = s.DATA_DB.Exec(
		`INSERT INTO health
            (suid, owner_uuid, sid, check_date, weight, length, topic, notes)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		suid, userUUID, req.Sid, req.CheckDate, req.Weight, req.Length, req.Topic, req.Notes,
	)
	if err != nil {
		http.Error(w, "error inserting health record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"health record added"}`))
}

//update health records

func (s *Server) snakeHealthUpdate(w http.ResponseWriter, r *http.Request) {
	// 1️⃣ Verify JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	// 2️⃣ Parse request body
	defer r.Body.Close()
	var req struct {
		Sid       string `json:"sid"`        // snake id (user-given)
		CheckDate string `json:"check_date"` // YYYY-MM-DD
		Weight    string `json:"weight,omitempty"`
		Length    string `json:"length,omitempty"`
		Topic     string `json:"topic,omitempty"`
		Notes     string `json:"notes,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.Sid == "" || req.CheckDate == "" {
		http.Error(w, "sid and check_date are required", http.StatusBadRequest)
		return
	}

	// Collect fields to update
	updates := []string{}
	args := []interface{}{}
	if req.Weight != "" {
		updates = append(updates, "weight = ?")
		args = append(args, req.Weight)
	}
	if req.Length != "" {
		updates = append(updates, "length = ?")
		args = append(args, req.Length)
	}
	if req.Topic != "" {
		updates = append(updates, "topic = ?")
		args = append(args, req.Topic)
	}
	if req.Notes != "" {
		updates = append(updates, "notes = ?")
		args = append(args, req.Notes)
	}

	if len(updates) == 0 {
		http.Error(w, "at least one field (weight, length, topic, or notes) is required", http.StatusBadRequest)
		return
	}

	// 3️⃣ Find suid for this snake
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

	// 4️⃣ Build dynamic UPDATE
	query := fmt.Sprintf(
		`UPDATE health SET %s WHERE suid = ? AND check_date = ?`,
		strings.Join(updates, ", "),
	)
	args = append(args, suid, req.CheckDate)

	res, err := s.DATA_DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "error updating health record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "health record not found", http.StatusNotFound)
		return
	}

	// 5️⃣ Respond success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"health record updated successfully"}`))
}

//delete health record

func (s *Server) snakeHealthDelete(w http.ResponseWriter, r *http.Request) {
	// 1️⃣ Verify JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	userUUID := claims["uuid"].(string)

	// 2️⃣ Parse request body
	defer r.Body.Close()
	var req struct {
		Sid       string `json:"sid"`        // user-provided snake ID
		CheckDate string `json:"check_date"` // YYYY-MM-DD of record to delete
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" || req.CheckDate == "" {
		http.Error(w, "sid and check_date are required", http.StatusBadRequest)
		return
	}

	// 3️⃣ Find suid for this snake
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

	// 4️⃣ Delete the health record
	res, err := s.DATA_DB.Exec(
		`DELETE FROM health WHERE suid = ? AND check_date = ?`,
		suid, req.CheckDate,
	)
	if err != nil {
		http.Error(w, "error deleting health record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "health record not found", http.StatusNotFound)
		return
	}

	// 5️⃣ Respond success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"health record deleted successfully"}`))
}
