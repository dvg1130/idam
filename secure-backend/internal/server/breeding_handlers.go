package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dvg1130/Portfolio/secure-backend/internal/auth"
	"github.com/dvg1130/Portfolio/secure-backend/internal/helpers"
	"github.com/dvg1130/Portfolio/secure-backend/models"
	datadb "github.com/dvg1130/Portfolio/secure-backend/repo/data_db"
)

// GET all SnakeBreeding events
func (s *Server) SnakeBreedGetAll(w http.ResponseWriter, r *http.Request) {
	//claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	//decode json

	defer r.Body.Close()

	req, err := helpers.DecodeBody[models.FindBreedingEvent](w, r)

	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" {
		http.Error(w, "sid is required", http.StatusBadRequest)
		return
	}

	//get suid

	var suid string
	err = s.Data_DB.QueryRow(
		datadb.GetSuid,
	).Scan(&suid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//get all breeding year values
	rows, err := s.Data_DB.Query(
		datadb.GetAllBreeding,
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

	//return list json

	resp := models.BreedingEventItem{
		Sid:           req.Sid,
		BreedingYears: years,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

// GET one SnakeBreed event
func (s *Server) SnakeBreedGetOne(w http.ResponseWriter, r *http.Request) {

}

// POST SnakeBreed event
func (s *Server) SnakeBreedPost(w http.ResponseWriter, r *http.Request) {
	//claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)

	//decode json
	defer r.Body.Close()

	req, err := helpers.DecodeBody[models.BreedingEvent](w, r)

	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.FemaleSid == "" || req.Male1Sid == "" || req.BreedingYear == "" {
		http.Error(w, "sid, mate_sid and breeding_year are required", http.StatusBadRequest)
		return
	}

	//get female suid
	var femaleSuid string
	err = s.Data_DB.QueryRow(
		datadb.GetSuid,
		ownerUUID, req.FemaleSid,
	).Scan(&femaleSuid)
	if err == sql.ErrNoRows {
		http.Error(w, "snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//get male(s) suid
	var male1Suid string
	var male2Suid string
	var male3Suid string
	var male4Suid string

	err = s.Data_DB.QueryRow(
		datadb.GetSuid,
		ownerUUID, req.Male1Sid,
	).Scan(&male1Suid)
	if err == sql.ErrNoRows {
		http.Error(w, "mate snake not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding mate snake: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if req.Male2Sid != "" {
		if err := s.Data_DB.QueryRow(datadb.GetSuid, ownerUUID, req.Male2Sid).Scan(&male2Suid); err != nil {
			http.Error(w, "error finding second mate: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if req.Male3Sid != "" {
		if err := s.Data_DB.QueryRow(datadb.GetSuid, ownerUUID, req.Male3Sid).Scan(&male3Suid); err != nil {
			http.Error(w, "error finding third mate: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if req.Male4Sid != "" {
		if err := s.Data_DB.QueryRow(datadb.GetSuid, ownerUUID, req.Male4Sid).Scan(&male4Suid); err != nil {
			http.Error(w, "error finding fourth mate: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	//add breeding event
	_, err = s.Data_DB.Exec(
		datadb.AddBreedingEvent,
		ownerUUID, femaleSuid, male1Suid, male2Suid, male3Suid, male4Suid,
		req.FemaleSid, req.Male1Sid, req.Male2Sid, req.Male3Sid, req.Male4Sid, req.BreedingYear, req.FemaleWeight,
		req.Male1Weight, req.Male2Weight, req.Male3Weight, req.Male4Weight,
		req.CoolingStart, req.CoolingEnd, req.WarmingStart, req.WarmingEnd, req.PairingDate1,
		req.PairingDate2, req.PairingDate3, req.PairingDate4, req.GravidDate, req.LayDate, req.ClutchSize,
		req.ClutchSurvive, req.Outcome, req.Notes,
	)
	if err != nil {
		http.Error(w, "error inserting breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"breeding record added"}`))

}

// UPDATE SnakeBreedevent
func (s *Server) SnakeBreedUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successful connection to SnakeBreedUpdate"))
}

// DELETE SnakeBreed event
func (s *Server) SnakeBreedDelete(w http.ResponseWriter, r *http.Request) {
	//claims
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return
	}
	ownerUUID := claims["uuid"].(string)
	//decode json
	defer r.Body.Close()

	req, err := helpers.DecodeBody[models.FindBreedingEvent](w, r)

	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.Sid == "" {
		http.Error(w, "sid is required", http.StatusBadRequest)
		return
	}

	// get suid
	var suid string
	err = s.Data_DB.QueryRow(
		datadb.GetSuid,
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
	err = s.Data_DB.QueryRow(
		datadb.GetBreedingUuid,
		ownerUUID, suid, req.BreedingYear,
	).Scan(&breedingUUID)
	if err == sql.ErrNoRows {
		http.Error(w, "breeding record not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "error finding breeding record: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
