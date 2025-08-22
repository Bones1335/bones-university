package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUserInternships(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserID      uuid.UUID `json:"user_id"`
		IntershipID uuid.UUID `json:"internship_id"`
	}

	type response struct {
		database.UserInternship
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
		respondWithError(w, http.StatusInternalServerError, "coudldn't decode user internship parameters", err)
		return
	}

	userInternship, err := cfg.db.CreateUserInternships(r.Context(), database.CreateUserInternshipsParams{
		UserID:       userID,
		InternshipID: params.IntershipID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user internship data", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{userInternship})
}
