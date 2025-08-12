package main

import (
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
)

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	type response struct {
		database.User
		UsersRole string
	}
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't get jwt", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't validate jwt", err)
		return
	}

	dbUser, err := cfg.db.GetSingleUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid UUID", err)
		return
	}

	usersRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		database.User{
			UsersID:         dbUser.UsersID,
			CreatedAt:       dbUser.CreatedAt,
			UpdatedAt:       dbUser.UpdatedAt,
			LastName:        dbUser.LastName,
			FirstName:       dbUser.FirstName,
			Username:        dbUser.Username,
			PersonalEmail:   dbUser.PersonalEmail,
			UniversityEmail: dbUser.UniversityEmail,
		},
		usersRole,
	})
}
