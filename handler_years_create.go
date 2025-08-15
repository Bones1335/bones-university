package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
)

func (cfg *apiConfig) handlerCreateYear(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		YearInSchool int16 `json:"year_in_school"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "coudln't validate JWT", err)
		return
	}

	userRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}

	if userRole != "administrator" {
		respondWithError(w, http.StatusUnauthorized, "user isn't authorized to complete this action", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get years parameters", err)
		return
	}

	year, err := cfg.db.CreateYears(r.Context(), params.YearInSchool)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create year in database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, year)
}
