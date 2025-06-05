package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"login_username"`
		Password string `json:"login_password"`
	}

	type response struct {
		database.User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect username or password", err)
		return
	}

	login, err := cfg.db.Login(r.Context(), params.Username)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "username not found", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, login.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect password", err)
		return
	}

	accessToken, err := auth.MakeJWT(login.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create access jwt", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create refresh token", err)
		return
	}

	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		UserID:    login.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't save refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: database.User{
			ID:              login.ID,
			CreatedAt:       login.CreatedAt,
			UpdatedAt:       login.UpdatedAt,
			PersonalEmail:   login.PersonalEmail,
			UniversityEmail: login.UniversityEmail,
			Username:        login.Username,
			LastName:        login.LastName,
			FirstName:       login.FirstName,
			Isstudent:       login.Isstudent,
			Isadmin:         login.Isadmin,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
