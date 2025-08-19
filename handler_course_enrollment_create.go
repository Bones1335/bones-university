package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateCourseEnrollment(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CourseID uuid.UUID `json:"course_id"`
	}

	type response struct {
		database.CourseEnrollment
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't validate JWT", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode course enrollment parameters", err)
		return
	}

	enrollment, err := cfg.db.CreateCourseEnrollment(r.Context(), database.CreateCourseEnrollmentParams{
		CourseID: params.CourseID,
		UserID:   userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create course enrollment", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{enrollment})
}
