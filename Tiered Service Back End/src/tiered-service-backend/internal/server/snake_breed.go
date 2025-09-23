package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tiered-service-backend/internal/auth"
)

//breeding handlers

//add breeding record

func (s *Server) snakeBreedAdd(w http.ResponseWriter, r *http.Request) {
	// ✅ Get user claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	defer r.Body.Close()
	// ✅ Parse request body
	var req struct {
		Sid           string  `json:"sid"`           // snake's ID
		MateSid       string  `json:"mate_sid"`      // mate's ID
		BreedingYear  string  `json:"breeding_year"` // YYYY-MM-DD or just YYYY-MM
		Weight        *string `json:"weight,omitempty"`
		CoolingStart  *string `json:"cooling_start,omitempty"`
		CoolingEnd    *string `json:"cooling_end,omitempty"`
		WarmingStart  *string `json:"warming_start,omitempty"`
		WarmingEnd    *string `json:"warming_end,omitempty"`
		PairingDate   *string `json:"pairing_date,omitempty"`
		GravidDate    *string `json:"gravid_date,omitempty"`
		LayDate       *string `json:"lay_date,omitempty"`
		ClutchSize    *int    `json:"clutch_size,omitempty"`
		ClutchSurvive *string `json:"clutch_survive,omitempty"`
		Outcome       *string `json:"outcome,omitempty"`
		Notes         *string `json:"notes,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// ✅ Validate required fields
	if req.Sid == "" || req.MateSid == "" || req.BreedingYear == "" {
		http.Error(w, "sid, mate_sid and breeding_year are required", http.StatusBadRequest)
		return
	}

	// ✅ Find primary snake suid
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Find mate suid
	var mateSuid string
	err = s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.MateSid,
	).Scan(&mateSuid)
	if err == sql.ErrNoRows {
		http.Error(w, "mate snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding mate snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Insert breeding record
	_, err = s.DATA_DB.Exec(`
		INSERT INTO breeding (
			owner_uuid, suid, mate_suid, sid, mate_sid,
			breeding_year, weight, cooling_start, cooling_end,
			warming_start, warming_end, pairing_date, gravid_date,
			lay_date, clutch_size, clutch_survive, outcome, notes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		ownerUUID, suid, mateSuid, req.Sid, req.MateSid,
		req.BreedingYear, req.Weight, req.CoolingStart, req.CoolingEnd,
		req.WarmingStart, req.WarmingEnd, req.PairingDate, req.GravidDate,
		req.LayDate, req.ClutchSize, req.ClutchSurvive, req.Outcome, req.Notes,
	)
	if err != nil {
		http.Error(w, "error inserting breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"breeding record added"}`))
}

// get breeeding record
func (s *Server) snakeBreedGetAll(w http.ResponseWriter, r *http.Request) {
	// ✅ Verify user authentication
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	defer r.Body.Close()

	// ✅ Parse request body for sid
	var req struct {
		Sid string `json:"sid"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" {
		http.Error(w, "sid is required", http.StatusBadRequest)
		return
	}

	// ✅ Find suid for the specified snake
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		// `SELECT * FROM breeding WHERE owner_uuid = ?AND breeding_year = ? AND (suid = ? OR mate_suid = ?)`//queires by either one of parent suids
		ownerUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Retrieve all breeding_year values for this snake
	rows, err := s.DATA_DB.Query(
		`SELECT breeding_year
		   FROM breeding
		  WHERE owner_uuid = ? AND suid = ?
		  ORDER BY breeding_year DESC`,
		ownerUUID, suid,
	)
	if err != nil {
		http.Error(w, "error retrieving breeding records: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	years := []string{}
	for rows.Next() {
		var year string
		if err := rows.Scan(&year); err != nil {
			http.Error(w, "error scanning record: "+err.Error(), http.StatusInternalServerError)
			return
		}
		years = append(years, year)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "row iteration error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Return list as JSON
	resp := struct {
		Sid           string   `json:"sid"`
		BreedingYears []string `json:"breeding_years"`
	}{
		Sid:           req.Sid,
		BreedingYears: years,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// get breeding single
func (s *Server) snakeBreedingGetOne(w http.ResponseWriter, r *http.Request) {
	// ✅ 1. Verify JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	defer r.Body.Close()

	// ✅ 2. Parse request body
	var req struct {
		Sid          string `json:"sid"`
		BreedingYear string `json:"breeding_year"` // expected YYYY format or YYYY-MM-DD if DATE
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" || req.BreedingYear == "" {
		http.Error(w, "sid and breeding_year are required", http.StatusBadRequest)
		return
	}

	// ✅ 3. Find the snake's suid
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ 4. Query the breeding record
	var record struct {
		Sid           string  `json:"sid"`
		MateSid       string  `json:"mate_sid"`
		BreedingYear  string  `json:"breeding_year"`
		Weight        *string `json:"weight,omitempty"`
		CoolingStart  *string `json:"cooling_start,omitempty"`
		CoolingEnd    *string `json:"cooling_end,omitempty"`
		WarmingStart  *string `json:"warming_start,omitempty"`
		WarmingEnd    *string `json:"warming_end,omitempty"`
		PairingDate   *string `json:"pairing_date,omitempty"`
		GravidDate    *string `json:"gravid_date,omitempty"`
		LayDate       *string `json:"lay_date,omitempty"`
		ClutchSize    *string `json:"clutch_size,omitempty"`
		ClutchSurvive *string `json:"clutch_survive,omitempty"`
		Outcome       *string `json:"outcome,omitempty"`
		Notes         *string `json:"notes,omitempty"`
	}

	query := `
		SELECT
			sid,
			mate_sid,
			breeding_year,
			weight,
			cooling_start,
			cooling_end,
			warming_start,
			warming_end,
			pairing_date,
			gravid_date,
			lay_date,
			clutch_size,
			clutch_survive,
			outcome,
			notes
		FROM breeding
		WHERE owner_uuid = ? AND suid = ? AND breeding_year = ?
	`

	err = s.DATA_DB.QueryRow(query, ownerUUID, suid, req.BreedingYear).Scan(
		&record.Sid,
		&record.MateSid,
		&record.BreedingYear,
		&record.Weight,
		&record.CoolingStart,
		&record.CoolingEnd,
		&record.WarmingStart,
		&record.WarmingEnd,
		&record.PairingDate,
		&record.GravidDate,
		&record.LayDate,
		&record.ClutchSize,
		&record.ClutchSurvive,
		&record.Outcome,
		&record.Notes,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "breeding record not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error retrieving record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ 5. Return JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(record); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

//update breeding record

func (s *Server) snakeBreedUpdate(w http.ResponseWriter, r *http.Request) {
	// ✅ Get user claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	defer r.Body.Close()

	// ✅ Parse request body
	var req struct {
		Sid           string  `json:"sid"`           // primary snake id (user-given)
		MateSid       string  `json:"mate_sid"`      // mate id (user-given)
		BreedingYear  string  `json:"breeding_year"` // required to locate record
		Weight        *string `json:"weight,omitempty"`
		CoolingStart  *string `json:"cooling_start,omitempty"`
		CoolingEnd    *string `json:"cooling_end,omitempty"`
		WarmingStart  *string `json:"warming_start,omitempty"`
		WarmingEnd    *string `json:"warming_end,omitempty"`
		PairingDate   *string `json:"pairing_date,omitempty"`
		GravidDate    *string `json:"gravid_date,omitempty"`
		LayDate       *string `json:"lay_date,omitempty"`
		ClutchSize    *int    `json:"clutch_size,omitempty"`
		ClutchSurvive *string `json:"clutch_survive,omitempty"`
		Outcome       *string `json:"outcome,omitempty"`
		Notes         *string `json:"notes,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// ✅ Validate required identifiers
	if req.Sid == "" || req.MateSid == "" || req.BreedingYear == "" {
		http.Error(w, "sid, mate_sid and breeding_year are required", http.StatusBadRequest)
		return
	}

	// ✅ Find suid for primary snake
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Find suid for mate snake
	var mateSuid string
	err = s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.MateSid,
	).Scan(&mateSuid)
	if err == sql.ErrNoRows {
		http.Error(w, "mate snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding mate snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ Build dynamic update list
	setParts := []string{}
	args := []any{}

	add := func(field string, val any) {
		setParts = append(setParts, field+" = ?")
		args = append(args, val)
	}

	if req.Weight != nil {
		add("weight", *req.Weight)
	}
	if req.CoolingStart != nil {
		add("cooling_start", *req.CoolingStart)
	}
	if req.CoolingEnd != nil {
		add("cooling_end", *req.CoolingEnd)
	}
	if req.WarmingStart != nil {
		add("warming_start", *req.WarmingStart)
	}
	if req.WarmingEnd != nil {
		add("warming_end", *req.WarmingEnd)
	}
	if req.PairingDate != nil {
		add("pairing_date", *req.PairingDate)
	}
	if req.GravidDate != nil {
		add("gravid_date", *req.GravidDate)
	}
	if req.LayDate != nil {
		add("lay_date", *req.LayDate)
	}
	if req.ClutchSize != nil {
		add("clutch_size", *req.ClutchSize)
	}
	if req.ClutchSurvive != nil {
		add("clutch_survive", *req.ClutchSurvive)
	}
	if req.Outcome != nil {
		add("outcome", *req.Outcome)
	}
	if req.Notes != nil {
		add("notes", *req.Notes)
	}

	// must have at least one field to update
	if len(setParts) == 0 {
		http.Error(w, "at least one field to update is required", http.StatusBadRequest)
		return
	}

	// ✅ Append identifiers to WHERE clause args
	args = append(args, ownerUUID, suid, mateSuid, req.BreedingYear)

	// ✅ Execute update
	query := fmt.Sprintf(`
		UPDATE breeding
		   SET %s
		 WHERE owner_uuid = ?
		   AND suid = ?
		   AND mate_suid = ?
		   AND breeding_year = ?`,
		strings.Join(setParts, ", "),
	)

	_, err = s.DATA_DB.Exec(query, args...)
	if err != nil {
		http.Error(w, "error updating breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"breeding record updated"}`))
}

// delete breeding record
func (s *Server) snakeBreedingDelete(w http.ResponseWriter, r *http.Request) {
	// ✅ 1. Verify JWT claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	defer r.Body.Close()

	// ✅ 2. Parse request body
	var req struct {
		Sid          string `json:"sid"`
		BreedingYear string `json:"breeding_year"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" || req.BreedingYear == "" {
		http.Error(w, "sid and breeding_year are required", http.StatusBadRequest)
		return
	}

	// ✅ 3. Find the snake's suid
	var suid string
	err := s.DATA_DB.QueryRow(
		`SELECT suid FROM snakes WHERE owner_uuid = ? AND sid = ?`,
		ownerUUID, req.Sid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ 4. Find breeding_uuid for this snake and breeding_year
	var breedingUUID string
	err = s.DATA_DB.QueryRow(
		`SELECT breeding_uuid
		   FROM breeding
		  WHERE owner_uuid = ? AND suid = ? AND breeding_year = ?`,
		ownerUUID, suid, req.BreedingYear,
	).Scan(&breedingUUID)
	if err == sql.ErrNoRows {
		http.Error(w, "breeding record not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ✅ 5. Delete the breeding record by breeding_uuid
	res, err := s.DATA_DB.Exec(
		`DELETE FROM breeding WHERE breeding_uuid = ?`,
		breedingUUID,
	)
	if err != nil {
		http.Error(w, "error deleting breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		http.Error(w, "breeding record not found", http.StatusNotFound)
		return
	}

	// ✅ 6. Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"breeding record deleted successfully"}`))
}

//update queries to use owner_uuid and a sid to find suid, then use suid and owner_uuid to find breeding_suid to find acount, so the request body can update mating/pariing
//and breeding yer if needed
