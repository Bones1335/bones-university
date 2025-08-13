package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerCreateDegrees(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		DegreeName       string `json:"degree_name"`
		DegreeLevel      string `json:"degree_level"`
		DegreeDepartment string `json:"degree_department"`
		DegreeDuration   int16  `json:"degree_duration"`
	}

	type response struct {
		database.DegreeProgram
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
	if userRole != "administrator" {
		respondWithError(w, http.StatusUnauthorized, "user isn't authorized to execute this action", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode degree program parameters", err)
		return
	}

	degree, err := cfg.db.CreateDegreePrograms(r.Context(), database.CreateDegreeProgramsParams{
		DegreeName:       params.DegreeName,
		DegreeLevel:      params.DegreeLevel,
		DegreeDepartment: params.DegreeDepartment,
		DegreeDuration:   params.DegreeDuration,
	})

	respondWithJSON(w, http.StatusOK, response{
		database.DegreeProgram{
			DegreesID:        degree.DegreesID,
			CreatedAt:        degree.CreatedAt,
			UpdatedAt:        degree.UpdatedAt,
			DegreeName:       degree.DegreeName,
			DegreeLevel:      degree.DegreeLevel,
			DegreeDepartment: degree.DegreeDepartment,
			DegreeDuration:   degree.DegreeDuration,
		},
	})
}
