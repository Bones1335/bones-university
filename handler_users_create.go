package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		PersonalEmail string `json:"personal_email"`
		LastName      string `json:"last_name"`
		FirstName     string `json:"first_name"`
		Password      string `json:"password"`
	}

	type response struct {
		database.User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't hash password", err)
	}

	username := generateUsername(params.FirstName, params.LastName)

	uniEmail := generateUniEmail(params.FirstName, params.LastName)

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		PersonalEmail:   params.PersonalEmail,
		LastName:        params.LastName,
		FirstName:       params.FirstName,
		Username:        username,
		UniversityEmail: uniEmail,
		Password:        hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		database.User{
			PersonalEmail:   user.PersonalEmail,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			Username:        user.Username,
			UniversityEmail: user.UniversityEmail,
		},
	})
}

func generateUsername(first, last string) string {
	fLow := strings.ToLower(first)
	lLow := strings.ToLower(last)

	beg := fLow[0]
	end := lLow[:5]

	return fmt.Sprintf("%s%s", string(beg), end)
}

func generateUniEmail(first, last string) string {
	fLow := strings.ToLower(first)
	lLow := strings.ToLower(last)

	return fmt.Sprintf("%s.%s@bones-uni.com", fLow, lLow)
}
