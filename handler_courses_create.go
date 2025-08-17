package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateCourses(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CourseCode        string    `json:"course_code"`
		CourseName        string    `json:"course_name"`
		CourseDescription string    `json:"course_description"`
		CourseProfessorID uuid.UUID `json:"course_professor_id"`
	}

	type response struct {
		database.Course
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
		respondWithError(w, http.StatusUnauthorized, "user isn't authorized to create courses", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode course parameters", err)
		return
	}

	course, err := cfg.db.CreateCourses(r.Context(), database.CreateCoursesParams{
		CourseCode:        params.CourseCode,
		CourseName:        params.CourseName,
		CourseDescription: params.CourseDescription,
		CourseProfessorID: params.CourseProfessorID,
	})

	respondWithJSON(w, http.StatusOK, response{
		course,
	})
}
