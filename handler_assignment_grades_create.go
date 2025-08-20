package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateAssignmentGrades(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		AssignmentID uuid.UUID `json:"assignment_id"`
		UserID       uuid.UUID `json:"user_id"`
	}

	type response struct {
		database.AssignmentGrade
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

	userRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}
	if userRole != "professor" && userRole != "administrator" {
		respondWithError(w, http.StatusUnauthorized, "user isn't authorized for this action", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode assignmnet grades parameters", err)
		return
	}

	assignmentGrades, err := cfg.db.CreateAssignmentGrades(r.Context(), database.CreateAssignmentGradesParams{
		AssignmentID: params.AssignmentID,
		UserID:       params.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create assignment grades", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{assignmentGrades})
}
