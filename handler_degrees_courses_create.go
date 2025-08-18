package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateDegreesCourses(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		DegreeID uuid.UUID `json:"degree_id"`
		CourseID uuid.UUID `json:"course_id"`
	}

	type response struct {
		database.DegreesCourse
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
		respondWithError(w, http.StatusInternalServerError, "couldn't decode degree's courses parameters", err)
		return
	}

	degreeCourse, err := cfg.db.CreateDegreesCourse(r.Context(), database.CreateDegreesCourseParams{
		DegreeID: params.DegreeID,
		CourseID: params.CourseID,
	})

	respondWithJSON(w, http.StatusOK, response{degreeCourse})
}
