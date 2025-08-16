package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateStudentsProgram(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CohortYear     int32     `json:"cohort_year"`
		StartDate      time.Time `json:"start_date"`
		StudentID      uuid.UUID `json:"student_id"`
		ProgramID      uuid.UUID `json:"program_id"`
		AcademicYearID uuid.UUID `json:"academic_year_id"`
	}

	type response struct {
		database.UsersProgram
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode students program parameters", err)
		return
	}

	studentProgram, err := cfg.db.CreateStudentsProgram(r.Context(), database.CreateStudentsProgramParams{
		CohortYear:     params.CohortYear,
		StartDate:      params.StartDate,
		StudentID:      params.StudentID,
		ProgramID:      params.ProgramID,
		AcademicYearID: params.AcademicYearID,
	})

	respondWithJSON(w, http.StatusOK, response{
		database.UsersProgram{
			UsersProgramID: studentProgram.UsersProgramID,
			CohortYear:     studentProgram.CohortYear,
			StartDate:      studentProgram.StartDate,
			StudentID:      studentProgram.StudentID,
			ProgramID:      studentProgram.ProgramID,
			AcademicYearID: studentProgram.AcademicYearID,
		},
	})
}
