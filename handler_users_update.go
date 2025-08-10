package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handleUpdateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LastName      string `json:"last_name"`
		FirstName     string `json:"first_name"`
		PersonalEmail string `json:"personal_email"`
		Password      string `json:"password"`
	}

	type response struct {
		database.User
		RoleName string
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't find JWT", err)
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
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't hash password", err)
		return
	}

	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		UsersID:       userID,
		LastName:      params.LastName,
		FirstName:     params.FirstName,
		PersonalEmail: params.PersonalEmail,
		Password:      hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update user", err)
		return
	}

	role, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		database.User{
			UsersID:         user.UsersID,
			UpdatedAt:       user.UpdatedAt,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			Username:        user.Username,
			PersonalEmail:   user.PersonalEmail,
			UniversityEmail: user.UniversityEmail,
		},
		role,
	})
}
