package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateAssignments(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		AssignmentName        string    `json:"assignment_name"`
		AssignmentDueDate     time.Time `json:"assignment_due_date"`
		AssignmentDescription string    `json:"assignment_description"`
		CourseID              uuid.UUID `json:"course_id"`
	}

	type response struct {
		database.Assignment
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

	usersRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}
	if usersRole != "professor" && usersRole != "administrator" {
		respondWithError(w, http.StatusUnauthorized, "user isn't authorized to execute this action", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode assignment parameters", err)
		return
	}

	assignment, err := cfg.db.CreateAssignments(r.Context(), database.CreateAssignmentsParams{
		AssignmentName:        params.AssignmentName,
		AssignmentDueDate:     params.AssignmentDueDate,
		AssignmentDescription: params.AssignmentDescription,
		CourseID:              params.CourseID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create assignment", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{assignment})
}
