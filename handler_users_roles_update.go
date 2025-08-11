package main

import (
	"encoding/json"
	"net/http"

	"github.com/Bones1335/bones-university/internal/auth"
	"github.com/Bones1335/bones-university/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUpdateUsersRole(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		RoleName string    `json:"role_name"`
		UsersID  uuid.UUID `json:"users_id"`
	}

	type response struct {
		database.UsersRole
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't validate jwt", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode role update parameters", err)
		return
	}

	userRole, err := cfg.db.GetUsersRole(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get user's role", err)
		return
	}

	switch userRole {
	case "administrator":
		role, err := cfg.db.GetSingleRole(r.Context(), params.RoleName)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "couldn't find role name for role change", err)
			return
		}

		newRole, err := cfg.db.UpdateUsersRole(r.Context(), database.UpdateUsersRoleParams{
			RoleID: role.RolesID,
			UserID: params.UsersID,
		})

		respondWithJSON(w, http.StatusOK, response{
			newRole,
		})
	case "professor", "student", "unset":
		respondWithError(w, http.StatusUnauthorized, "user isn't an administrator and isn't authorized to change a user role", err)
		return
	default:
		respondWithError(w, http.StatusUnauthorized, "the user's role doesn't exist and therefore isn't authorized", err)
		return
	}
}
